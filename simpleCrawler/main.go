package main

import (
	"fmt"
	"regexp"

	"simpleCrawler/common"
	"simpleCrawler/engine"
	. "simpleCrawler/love/parser"
)

func main() {

	engine.Run(common.Request{
		Url:        "http://www.zhenai.com/zhenhun",
		ParserFunc: ParserCityList,
	})
	return
	// req := new(common.Request)
	// req.Url = "d"
	// resp, err := http.Get("http://www.zhenai.com/zhenhun")

	// if err != nil {
	// 	panic(err)
	// }

	// defer resp.Body.Close()

	// if resp.StatusCode == http.StatusOK {

	// 	ed := determineEncoding(resp.Body)
	// 	uReader := transform.NewReader(resp.Body, ed.NewDecoder())
	// 	bd, err := ioutil.ReadAll(uReader)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Printf("%s\n", bd)
	// 	ParserCityList(bd)
	// } else {
	// 	fmt.Printf("error :", resp.StatusCode)
	// }
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
