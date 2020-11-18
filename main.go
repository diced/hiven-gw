package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/diced/hivengw/gateway"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	gate := gateway.NewGateway(gateway.ParseEnv())

	if gate.Config.CompressionZlib {
		log.Println("using zlib compression method")
	}

	for {
		var msg gateway.Map
		_, m, err := gate.Websocket.Socket.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		if gate.Config.CompressionZlib {
			gate.Websocket.Uncompress(m, &msg)
		} else {
			json.Unmarshal(m, &msg)
		}
		
		if !gateway.CheckEmpty("DEBUG") {
			log.Println("op:", msg["op"], " e:", msg["e"])
		}

		switch msg["op"] {
		case float64(1):
			var overall [][]int
			go func() {
				for {
					gate.Websocket.Heartbeat()
					if !gateway.CheckEmpty("DEBUG") {
						a := make([]int, 0, 999999)
						overall = append(overall, a)
						gate.Stats(true)
						overall = nil
					}
					time.Sleep(30 * time.Second)
				}
			}()
			gate.Websocket.Reconnect(gate.Config.Token)
		default:
			b, err := json.Marshal(msg)
			if err != nil {
				log.Fatal(err)
			}
			_, err = gate.Redis.Do("RPUSH", gate.Config.List, string(b))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
