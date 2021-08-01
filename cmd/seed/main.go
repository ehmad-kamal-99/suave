package main

import (
	"os"

	"github.com/brianvoe/gofakeit/v6"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/ehmad-kamal-99/suave/config"
	"github.com/ehmad-kamal-99/suave/db/redisbloom"
)

func main() {
	gofakeit.Seed(0)
	usernameList := make([]string, 0) // list to store 10000 fake random usernames

	_, err := os.Stat("seed.txt")
	if err == nil {
		log().Debugf("deleting old seed file found.")
		_ = os.Remove("seed.txt")
	} else if os.IsNotExist(err) {
		log().Debugf("creating seed file to check usernames.")
	}

	file, err := os.OpenFile("seed.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// generate 10000 random usernames
	for i := 0; i < 10000; i++ {
		usrName := gofakeit.Username()
		usernameList = append(usernameList, usrName)
		_, err = file.Write([]byte(usrName + "\n"))
		if err != nil {
			log().Panicf("failed to write username to file, err: %+v", err)
			panic(err)
		}
	}

	// new redis bloom client
	cli := redisbloom.NewClient()

	// add usernames in redis-bloomfilter
	log().Info("adding 10000 usernames in suave-db [redis-bloomfilter]")
	if err := cli.AddUsernames(usernameList); err != nil {
		panic(err)
	}

	log().Info("added fake usernames in redis-bloomfilter")
	cli.Close()
}

func log() *logger.Entry {
	level, err := logger.ParseLevel(viper.GetString(config.LogLevel))
	if err != nil {
		logger.SetLevel(logger.DebugLevel)
	}
	logger.SetLevel(level)

	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp: true,
	})

	return logger.WithFields(logger.Fields{
		"package": "seed-main",
	})
}
