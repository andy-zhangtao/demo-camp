package main

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

func main() {
	// no need auth
	client := influxdb2.NewClient("http://xxxxx:8086/", "")
	health, err := client.Health(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v \n", health)

	writeAPI := client.WriteApiBlocking("", "test/autogen")
	// create point using full params constructor
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())
	// Write data
	err = writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		fmt.Printf("Write error: [%s]\n", err.Error())
	}

	queryApi := client.QueryApi("")
	// Supply string in a form database/retention-policy as a bucket. Skip retention policy for the default one, use just a database name (without the slash character)
	result, err := queryApi.Query(context.Background(), `from(bucket:"test")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
	if err == nil {
		for result.Next() {
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			fmt.Printf("row: %s\n", result.Record().String())
		}
		if result.Err() != nil {
			fmt.Printf("Query error: %s\n", result.Err().Error())
		}
	} else {
		fmt.Printf("Query error: %s\n", err.Error())
	}
	// Close client
	client.Close()
}
