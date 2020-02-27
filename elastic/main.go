package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/olivere/elastic/v7"
)

type NginxLog struct {
	Domain               string    `json:"domain"`
	RemoteAddr           string    `json:"remote_addr"`
	RemoteUser           string    `json:"remote_user"`
	TimeLocal            string    `json:"time_local"`
	Request              string    `json:"request"`
	Status               string    `json:"status"`
	BodyBytesSent        string    `json:"body_bytes_sent"`
	HTTPReferer          string    `json:"http_referer"`
	HTTPUserAgent        string    `json:"http_user_agent"`
	HTTPXForwardedFor    string    `json:"http_x_forwarded_for"`
	CookieJSESSIONID     string    `json:"cookie_JSESSIONID"`
	UpstreamAddr         string    `json:"upstream_addr"`
	UpstreamStatus       string    `json:"upstream_status"`
	UpstreamResponseTime string    `json:"upstream_response_time"`
	RequestTime          string    `json:"request_time"`
	RequestBody          string    `json:"request_body"`
	CookieVer            string    `json:"cookie__ver"`
	Timestamp            time.Time `json:"@timestamp"`
}

func main() {
	ctx := context.Background()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c := &http.Client{Transport: tr}

	client, err := elastic.NewClient(
		elastic.SetHttpClient(c),
		elastic.SetURL("https://192.168.2.188:9200"),
		elastic.SetScheme("https"),
		elastic.SetBasicAuth("xxxx", "xxxx"),
	)

	if err != nil {
		// Handle error
		panic(err)
	}

	info, code, err := client.Ping("https://192.168.2.188:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	termQuery := elastic.NewTermQuery("domain", "postinoa.eqshow.cn")
	searchResult, err := client.Search().
		Index("testenv_nginx*").   // search in index "twitter"
		Query(termQuery).          // specify the query
		Sort("@timestamp", false). // sort by "user" field, ascending
		From(0).Size(10).          // take documents 0-9
		Pretty(true).              // pretty print request and response JSON
		Do(ctx)                    // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	if searchResult.Hits.TotalHits.Value > 0 {

		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var nl NginxLog
			err := json.Unmarshal(*&hit.Source, &nl)
			if err != nil {
				// Deserialization failed
			}

			// Work with tweet
			// nl.Timestamp.Format("2006-01-02T15:04:05")
			fmt.Printf("%s %s %s %s \n", nl.TimeLocal, nl.Domain, nl.Request, nl.Status)
		}
	} else {
		// No hits
		fmt.Print("Found no tweets\n")
	}

	fmt.Println("range demo")

	rangeDemoWithFields(client)
}
