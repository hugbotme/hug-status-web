package config

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"log"
)

type Configuration struct {
	Redis     RedisConfiguration     `json:"redis"`
	StatusWeb StatusWebConfiguration `json:"status-web"`
}

type StatusWebConfiguration struct {
	URL string `json:"url"`
}

type RedisConfiguration struct {
	URL  string `json:"url"`
	Auth string `json:"auth"`
}

func NewConfiguration(configFile *string) (*Configuration, error) {
	fileContent, err := ioutil.ReadFile(*configFile)
	if err != nil {
		return nil, err
	}

	var config Configuration
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (config *Configuration) ConnectRedis() redis.Conn {
	redisClient, err := redis.Dial("tcp", config.Redis.URL)
	if err != nil {
		log.Fatal("Redis client init (connect) failed:", err)
	}

	if len(config.Redis.Auth) == 0 {
		return redisClient
	}

	if _, err := redisClient.Do("AUTH", config.Redis.Auth); err != nil {
		redisClient.Close()
		log.Fatal("Redis client init (auth) failed:", err)
	}

	return redisClient
}
