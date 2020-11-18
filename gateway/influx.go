package gateway

import influxdb2 "github.com/influxdata/influxdb-client-go/v2"

// Influx struct for holding stuff about influx?
type Influx struct {
	Host   string
	Bucket string
	Org    string
	Token  string
	Client influxdb2.Client
}

// NewInflux creates a new instance of Influx
func NewInflux(host string, bucket string, org string, token string) Influx {
	client := influxdb2.NewClient(host, token)

	return Influx{
		Host:   host,
		Bucket: bucket,
		Org:    org,
		Token:  token,
		Client: client,
	}
}
