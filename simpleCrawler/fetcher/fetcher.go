package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	//. "simpleCrawler/love/parser"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return nil, fmt.Errorf("wrong status code %v", resp.StatusCode)

		// fmt.Printf("%s\n", bd)
		// outstring = ParserCityList(bd)
	}

	ed := determineEncoding(resp.Body)
	uReader := transform.NewReader(resp.Body, ed.NewDecoder())
	return ioutil.ReadAll(uReader)

}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)

	if err != nil {
		panic(err)
	}

	e, _, certain := charset.DetermineEncoding(bytes, "")

	if certain == true {
		return e
	}

	return e
	//return unicode.UTF8
}
