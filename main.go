package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/diced/hivengw/gateway"
	"github.com/joho/godotenv"
)

func enabledEvent(events []string, event string) bool {
	for i := range events {
		if events[i] == event {
			return false
		}
	}

	return true
}

func main() {
	godotenv.Load()

	gate := gateway.NewGateway(gateway.ParseEnv())
	log.Println(gate.Config.DisabledEvents)

	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt)

	if gate.Config.CompressionZlib {
		log.Println("using zlib compression method")
	}

	// if gate.Config.Influx.Client != nil {
	// 	// to be implemented
	// }

	for {
		var msg gateway.HivenResponse
		_, m, err := gate.Websocket.Socket.ReadMessage()
		if err != nil {
			log.Fatal(m, err)
		}

		if gate.Config.CompressionZlib {
			gate.Websocket.Uncompress(m, &msg)
		} else {
			_, r, err := gate.Websocket.Socket.NextReader()
			if err != nil {
				log.Fatal(err)
			}
			err = json.NewDecoder(r).Decode(&msg)
			if err != nil {
				log.Fatal(err)
			}
		}

		if enabledEvent(gate.Config.DisabledEvents, msg.Event) {
			if !gateway.CheckEmpty("DEBUG") {
				log.Println("op:", msg.OpCode, " e:", msg.Event)
			}

			switch msg.OpCode {
			case 1:
				var overall [][]int
				done := make(chan bool)
				go func() {
					defer close(done)
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
}
