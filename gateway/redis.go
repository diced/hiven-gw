package gateway

// RedisCmd struct
type RedisCmd struct {
	OpCode int                    `json:"op"`
	Data   map[string]interface{} `json:"data"`
}