package parser

import (
	"fmt"
	"simpleCrawler/fetcher"
	"testing"
)

const UrlTest = "http://www.zhenai.com/zhenhun"

func TestParserCityList(t *testing.T) {
	contents, err := fetcher.Fetch(UrlTest)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", contents)

}
