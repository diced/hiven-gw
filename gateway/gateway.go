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

	encoding := "zlib_json"

	if !config.CompressionZlib {
		encoding = "text_json"
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
	b, err := json.Marshal(Map{"e": "stats", "stats": m})
	if err != nil {
		log.Println("could not marshal stats skipping")
	} else {
		_, err = g.Redis.Do("RPUSH", g.Config.List, string(b))
		if err != nil {
			log.Println("could not rpush stats")
		}
	}
	log.Printf("%valloc: %v mb  totalloc: %v mb  sys: %v mb  gc: %v\n", s, m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024, m.NumGC)
}

// DebugLog sends a debuglog with hiven res
func (g *Gateway) DebugLog(msg HivenResponse) {
	if !CheckEmpty("DEBUG") {
		log.Println("op:", msg.OpCode, " e:", msg.Event)
	}
}
