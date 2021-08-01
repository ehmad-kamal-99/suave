package health

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/viper"

	"github.com/ehmad-kamal-99/suave/config"
)

// constants for service state and connection timout.
const (
	SafeState      = "SAFE"
	WarnState      = "WARN"
	UnhealthyState = "UNHEALTHY"
	OKState        = "OK"

	TimeoutSeconds = 3
)

// Status - struct representing service health status.
type Status struct {
	State      string `json:"state"`
	Main       string `json:"main"`
	RedisBloom string `json:"mongo"`
}

// Health - interface holding implementation for getting service's health status.
type Health interface {
	GetSvcHealth() Status
}

type service struct {
	state         string
	mainURL       string
	redisBloomURL string
}

// New - initiates new service struct.
func New() Health {
	return &service{
		state:         SafeState,
		mainURL:       fmt.Sprintf("%s:%s", viper.GetString(config.ServerHost), viper.GetString(config.ServerPort)),
		redisBloomURL: fmt.Sprintf("%s:%s", viper.GetString(config.DBHost), viper.GetString(config.DBPort)),
	}
}

// GetSvcHealth - checks all service's health status.
func (s *service) GetSvcHealth() Status {
	mainStatus := OKState
	if status := checkHealth(s.mainURL); !status {
		mainStatus = "Main Error: Application Node not responding"
	}

	redisBloomStatus := OKState
	if status := checkHealth(s.redisBloomURL); !status {
		redisBloomStatus = "RedisBloom Error: Database Node not responding"
	}

	switch {
	case mainStatus == OKState && redisBloomStatus == OKState:
		s.state = SafeState
	case mainStatus != OKState && redisBloomStatus != OKState:
		s.state = UnhealthyState
	default:
		s.state = WarnState
	}

	return Status{
		State:      s.state,
		Main:       mainStatus,
		RedisBloom: redisBloomStatus,
	}
}

// checkHealth - tries connecting to node/service and returns response as boolean.
func checkHealth(address string) bool {
	conn, err := net.DialTimeout("tcp", address, TimeoutSeconds*time.Second)
	if err != nil {
		return false
	}

	if conn != nil {
		_ = conn.Close()
		return true
	}

	return false
}
