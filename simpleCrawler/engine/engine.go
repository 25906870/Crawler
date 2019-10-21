package engine

import (
	"io/ioutil"
	"log"
	. "simpleCrawler/common"
	"simpleCrawler/fetcher"
	"time"
)

func Wloghtml(ctx []byte) {
	//time.Sleep(time.Second)
	filename := time.Now()
	// rand.Seed(time.Now().UnixNano())
	// sec := rand.Intn(1000)

	filestr := filename.Format("2006-01-02_15_04_05.000000") + ".html"
	err := ioutil.WriteFile("log/"+filestr, ctx, 0644)
	if err != nil {
		log.Printf("wloghtml failed %v", err)
	}
}
func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {

		r := requests[0]
		requests = requests[1:]

		body, err := fetcher.Fetch(r.Url)

		if err != nil {
			log.Printf("Fetcher : error "+
				"fetching url %s: %v", r.Url, err)
		}
		//fmt.Printf("%v", string(body))
		Wloghtml(body)
		return
		ParseResult := r.ParserFunc(body)
		requests = append(requests, ParseResult.Request...)

		for _, item := range ParseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
