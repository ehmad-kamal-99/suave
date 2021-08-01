package redisbloom

import (
	"fmt"

	rdsbloom "github.com/RedisBloom/redisbloom-go"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/ehmad-kamal-99/suave/config"
	"github.com/ehmad-kamal-99/suave/db"
)

// client holds redisbloom client struct.
type client struct {
	cli *rdsbloom.Client
}

// NewClient - initiates a new redisbloom client.
func NewClient() db.Datastore {
	log().Infof("initializing redis connection at %s:%s", viper.GetString(config.DBHost), viper.GetString(config.DBPort))

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%s", viper.GetString(config.DBHost), viper.GetString(config.DBPort)), redis.DialPassword(viper.GetString(config.DBPassword)))
		},
	}

	redisClient := rdsbloom.NewClientFromPool(pool, viper.GetString(config.DBClientName))

	return &client{cli: redisClient}
}

// AddUsername - add username into database.
func (c *client) AddUsername(username string) error {
	if _, err := c.cli.Add(viper.GetString(config.DBFilterName), username); err != nil {
		log().Errorf("failed to add username: %s, err: %+v", username, err)
		return errors.Wrap(err, "failed to add username in database")
	}

	return nil
}

// CheckUsername - checks username in database.
func (c *client) CheckUsername(username string) (bool, error) {
	exists, err := c.cli.Exists(viper.GetString(config.DBFilterName), username)
	if err != nil {
		log().Errorf("failed to check for username: %s, err: %+v", username, err)
		return false, errors.Wrap(err, "failed to check username in database")
	}

	return exists, nil
}

// Close - closed redisbloom client connection.
func (c *client) Close() error {
	if err := c.cli.Pool.Close(); err != nil {
		log().Errorf("failed to close redis connection, err: %+v", err)
		return errors.Wrap(err, "failed to close database connection")
	}

	return nil
}
