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

	zlib := true

	if CheckEmpty("ZLIB") {
		zlib = false
	}

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
	}
}

// CheckEmpty checks if a env var exists
func CheckEmpty(v string) bool {
	env := os.Getenv(v)

	if len(strings.TrimSpace(env)) == 0 {
		return true
	}
	return false
}
