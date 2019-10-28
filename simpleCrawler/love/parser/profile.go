package parser

import (
	"encoding/json"
	"fmt"
	"regexp"
	"simpleCrawler/common"
	"strconv"
)

const ItemRe = `<div class="list-item">.*<div class="item-btn">[^<]*</div></div>`

var linknameRe = regexp.MustCompile(`<th><a href="http://album.zhenai.com/u/([^"]+)" [^>]*>([^<]*)</a></th>`)

var ageRe = regexp.MustCompile(`<td[^>]*><span [^>]*>年龄：</span>([\d]+)</td>`)
var genderRe = regexp.MustCompile(`<td [^>]*><span [^>]*>性别：</span>([^<]+)</td>`)
var marrRe = regexp.MustCompile(`<td[^>]*><span [^>]*>婚况：</span>([^<]+)</td>`)
var HeightRe = regexp.MustCompile(`<td [^>]*><span [^>]*>身[^高]*高：</span>([\d]+)</td>`)
var edu = regexp.MustCompile(`<td [^>]*><span [^>]*>学[^历]*历：</span>([^<]+)</td>`)

const ItemCount = 3

func ParserProfile(contents []byte) common.ParseResult {

	result := common.ParseResult{}
	itRegxp := regexp.MustCompile(ItemRe)
	itemlist := itRegxp.FindSubmatch(contents)
	itemlistlength := len(itemlist)

	for index := 0; index < ItemCount && itemlistlength > index; index++ {
		prf := common.Profile{}

		link_name := linknameRe.FindSubmatch(itemlist[index])

		if len(link_name) > 1 {
			prf.Name = string(link_name[0]) + "_" + string(link_name[1])
		}

		age, err := strconv.Atoi(submatch(itemlist[index], ageRe))
		if err == nil {
			prf.Age = age
		} else {
			fmt.Printf("submatch %v", err)
		}

		prf.Gender = submatch(itemlist[index], genderRe)
		prf.Marriage = submatch(itemlist[index], marrRe)
		prf.Education = submatch(itemlist[index], edu)
		prf.Height, err = strconv.Atoi(submatch(itemlist[index], HeightRe))
		if err != nil {
			prf.Height = 100
		}
		bt, err := json.Marshal(prf)

		if err == nil {
			common.Wlogfile(prf.Name, bt)
		}
		result.Items = append(result.Items, prf)
	}
	return result
}

func submatch(ctx []byte, reg *regexp.Regexp) string {

	matchs := reg.FindSubmatch(ctx)

	if len(matchs) > 1 {
		return string(matchs[1])
	}

	return ""
}
