package common

import (
	"io/ioutil"
	"log"
	"time"
)

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hokou      string
	House      string
	Car        string
}

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Request []Request
	Items   []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

func Wloghtml(ctx []byte) {
	return
	filename := time.Now()

	filestr := filename.Format("2006-01-02_15_04_05.000000") + ".html"
	err := ioutil.WriteFile("log/"+filestr, ctx, 0644)
	if err != nil {
		log.Printf("wloghtml failed %v", err)
	}
}

func Wlogfile(filename string, ctx []byte) {

	err := ioutil.WriteFile("log/"+filename, ctx, 0644)
	if err != nil {
		log.Printf("wloghtml failed %v", err)
	}
}
