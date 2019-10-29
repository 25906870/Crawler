package main

import (
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
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

	common.Wloghtml(body)
	return true
}

func main() {

	// go func() {
	// 	http.ListenAndServe("localhost:12351", nil)
	// }()
	engine.Run(common.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: ParserCityList,
	})

	return

}
