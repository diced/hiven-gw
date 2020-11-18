package gateway

import (
	"log"
	"os"
	"strings"
)

// Env config vars
type Env struct {
	Token           string
	Redis           string
	List            string
	DisabledEvents  []string
	Influx          Influx
	CompressionZlib bool
}

// ParseEnv from dotenv
func ParseEnv() Env {
	envs := []string{"TOKEN", "REDIS", "LIST"}

	for _, k := range envs {
		if CheckEmpty(k) {
			log.Fatalf("Env %v was required but not found in the environ...\n", k)
		}
	}

	zlib := false

	if !CheckEmpty("ZLIB") {
		zlib = true
	}

	influx := InfluxConfig()
	var disabledEv []string
	if !CheckEmpty("DISABLED_EVENTS") {
		disabledEv = strings.Split(os.Getenv("DISABLED_EVENTS"), ",")
	}

	return Env{
		Token:           os.Getenv("TOKEN"),
		Redis:           os.Getenv("REDIS"),
		List:            os.Getenv("LIST"),
		DisabledEvents:  disabledEv,
		CompressionZlib: zlib,
		Influx:          influx,
	}
}

// InfluxConfig checks for influx env's and returns a influx config
func InfluxConfig() Influx {
	envs := []string{"INFLUX_TOKEN", "INFLUX_ORG", "INFLUX_BUCKET", "INFLUX_HOST"}

	for _, k := range envs {
		if CheckEmpty(k) {
			return Influx{
				Bucket: "",
				Client: nil,
				Host:   "",
				Org:    "",
				Token:  "",
			}
		}
	}

	return NewInflux(os.Getenv("INFLUX_HOST"), os.Getenv("INFLUX_BUCKET"), os.Getenv("INFLUX_ORG"), os.Getenv("INFLUX_TOKEN"))
}

// CheckEmpty checks if a env var exists
func CheckEmpty(v string) bool {
	env := os.Getenv(v)

	if len(strings.TrimSpace(env)) == 0 {
		return true
	}
	return false
}
