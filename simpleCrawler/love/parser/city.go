package parser

import (
	"fmt"
	"regexp"
	"simpleCrawler/common"
)

const CityReg = `<a [^h]*href="(http://[a-zA-Z]{1,5}.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) common.ParseResult {
	re := regexp.MustCompile(CityReg)

	matches := re.FindAllSubmatch(contents, -1)
	result := common.ParseResult{}
	for _, m := range matches {
		//for _, sbm := range m
		nm := string(m[2])
		fmt.Printf("City:%5s ,URL: %s \n", m[2], m[1])
		result.Items = append(result.Items, nm)
		result.Request = append(result.Request, common.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) common.ParseResult {
				return ParserProfile(c, nm)
			},
		})
		//outstr = append(outstr, string(m[1]))

	}
	fmt.Printf("Matches Nums %d\n", len(matches))
	return result
}
