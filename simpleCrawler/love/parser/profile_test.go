package parser

import (
	"simpleCrawler/common"
	"io/ioutil"
	"testing"
)

const Testfile = "profiletest.html"

func TestParserProfile(t *testing.T) {

	filedt, err := ioutil.ReadFile(Testfile)

	if err != nil {
		panic(err)
	}

	common.ParseResult := ParserProfile(filedt)

	count := len(common.ParseResult.Requests)

	for index := 0; index < count; index++ {
		
	}
}
