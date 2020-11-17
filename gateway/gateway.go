package gateway

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

// Gateway struct
type Gateway struct {
	Redis     redis.Conn
	Websocket Websocket
	Config    Env
}

// NewGateway creates a new gateway struct with redis
func NewGateway(config Env) Gateway {
	redis, err := redis.Dial("tcp", config.Redis)
	if err != nil {
		log.Fatal(err)
	}

	return Gateway{
		Redis:     redis,
		Websocket: NewWebsocket("wss://swarm-dev.hiven.io/socket?encoding=json&compression=textjson"),
		Config:    config,
	}
}
