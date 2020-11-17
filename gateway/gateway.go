package gateway

import (
	"log"
	"runtime"

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
		Websocket: NewWebsocket("wss://swarm-dev.hiven.io/socket?encoding=json&compression=text_zlib"),
		Config:    config,
	}
}

// Stats gets memory stats...
func (g *Gateway) Stats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("alloc: %v mb  totalloc: %v mb  sys: %v mb  gc: %v\n", m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024, m.NumGC)
}
