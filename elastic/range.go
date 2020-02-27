package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"
)

func rangeDemoWithFields(client *elastic.Client) {
	statusBoolQuery := elastic.NewBoolQuery()
	dateFilter := elastic.NewBoolQuery()
	dateFilter = dateFilter.Filter(elastic.NewRangeQuery("@timestamp").Gte("now-3m").Lte("now"))

	statusBoolQuery.Must(elastic.NewMatchQuery("request", "/v3"), dateFilter)

	searchResult, err := client.Search().
		Index("testenv_nginx*").   // search in index "twitter"
		Query(statusBoolQuery).    // specify the query
		Sort("@timestamp", false). // sort by "user" field, ascending
		From(0).Size(50).          // take documents 0-9
		Pretty(true).              // pretty print request and response JSON
		Do(context.Background())   // execute
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
}

func rangeDemo(client *elastic.Client) {
	//rangeQuery := elastic.NewRangeQuery("@timestamp").From("now - 1d").To("now")

	dateFilter := elastic.NewBoolQuery()
	dateFilter = dateFilter.Filter(elastic.NewRangeQuery("@timestamp").Gte("now-3m").Lte("now"))

	searchResult, err := client.Search().
		Index("testenv_nginx*").   // search in index "twitter"
		Query(dateFilter).         // specify the query
		Sort("@timestamp", false). // sort by "user" field, ascending
		From(0).Size(50).          // take documents 0-9
		Pretty(true).              // pretty print request and response JSON
		Do(context.Background())   // execute
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
}
