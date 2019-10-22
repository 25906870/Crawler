package parser

import (
	"io/ioutil"
	"testing"
)

const UrlTest = "http://www.zhenai.com/zhenhun"

func TestParserCityList(t *testing.T) {

	ctn, err := ioutil.ReadFile("parserTest.html")

	if err != nil {
		panic(err)
	}

	resluts := ParserCityList(ctn)

	const resultsize = 37
	size := len(resluts.Request)
	if size == resultsize {
		t.Errorf("result should have %v"+"actully %v", resultsize, len(resluts.Items))
	}
	t.Log("TEST END")
}
