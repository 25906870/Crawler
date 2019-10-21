package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"simpleCrawler/common"
	"simpleCrawler/engine"
	. "simpleCrawler/love/parser"
)

func PostMantest() bool {
	url := "http://www.zhenai.com/zhenghun"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "de188d59-71db-8e67-e376-139eb3fade0f")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	engine.Wloghtml(body)
	return true
}

func main() {

	engine.Run(common.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: ParserCityList,
	})
	return

}

func processCityList(contents []byte) {
	re := regexp.MustCompile(`<a \S* href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

	matches := re.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		//for _, sbm := range m
		{
			fmt.Printf("City:%5s ,URL: %s \n", m[2], m[1])
		}

	}
	fmt.Printf("Matches Nums %d\n", len(matches))
}
