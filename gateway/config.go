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

	return Env{
		Token:           os.Getenv("TOKEN"),
		Redis:           os.Getenv("REDIS"),
		List:            os.Getenv("LIST"),
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
