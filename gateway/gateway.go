package gateway

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"

	"github.com/gomodule/redigo/redis"
)

// Map alias for map[string]interface{}
type Map map[string]interface{}

// HivenResponse swarm response
type HivenResponse struct {
	OpCode int         `json:"op"`
	Data   interface{} `json:"d"`
	Event  string      `json:"e"`
	Seq    int         `json:"seq"`
}

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

	encoding := "text_json"

	if config.CompressionZlib {
		encoding = "zlib_json"
	}

	return Gateway{
		Redis:     redis,
		Websocket: NewWebsocket(fmt.Sprintf("wss://swarm-dev.hiven.io/socket?encoding=json&compression=%v", encoding)),
		Config:    config,
	}
}

// Stats gets memory stats...
func (g *Gateway) Stats(hb bool) {
	s := "hb -> "
	if !hb {
		s = ""
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	b, err := json.Marshal(m)
	if err != nil {
		log.Println("could not marshall stats, skipping sending")
	} else {
		g.Redis.Do("RPUSH", "stats", b)
	}

	log.Printf("%valloc: %v mb  totalloc: %v mb  sys: %v mb  gc: %v\n", s, m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024, m.NumGC)
}