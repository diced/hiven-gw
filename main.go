package main

import (
	"encoding/json"
	"log"

	"github.com/diced/hivengw/gateway"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	gate := gateway.NewGateway(gateway.ParseEnv())

	for {
		var msg map[string]interface{}
		err := gate.Websocket.Socket.ReadJSON(&msg)
		if err != nil {
			log.Fatal(err)
		}

		if !gateway.CheckEmpty("DEBUG") {
			log.Println("op:", msg["op"], " e:", msg["e"])
		}

		switch msg["op"] {
		case float64(1):
			go gate.Websocket.Heartbeat()
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
