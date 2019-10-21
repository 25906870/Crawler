package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	//. "simpleCrawler/love/parser"

	"github.com/pborman/uuid"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	time.Sleep(time.Second)
	req, _ := http.NewRequest("GET", url, nil)

	u1 := uuid.NewUUID()
	uustr := fmt.Sprintf("%s", u1)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", uustr)
	resp, err := http.DefaultClient.Do(req)

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
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
