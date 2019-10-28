package engine

import (
	"log"
	. "simpleCrawler/common"
	"simpleCrawler/fetcher"
	"time"
)

func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}
	for {
		for len(requests) > 0 {

			r := requests[0]
			requests = requests[1:]

			body, err := fetcher.Fetch(r.Url)

			if err != nil {
				log.Printf("Fetcher : error "+
					"fetching url %s: %v", r.Url, err)
			}
			//fmt.Printf("%v", string(body))
			if len(body) > 0 {
				Wloghtml(body)
			}

			ParseResult := r.ParserFunc(body)
			requests = append(requests, ParseResult.Request...)

			// for _, item := range ParseResult.Items {
			// 	log.Printf("Got item %v", item)
			// }
		}
		time.Sleep(time.Second)
	}

}
