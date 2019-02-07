package parser

import (
	"regexp"

	"github.com/GoSpider/chanCrawler/model"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhanghun/[0-9a-z]+)"[^>]*>([^<]*)</a>`

func ParserCityList(contents []byte) model.ParserResult {
	result := model.ParserResult{}
	// 正则匹配() 用于提取
	rg := regexp.MustCompile(cityListRe)
	allSubmatch := rg.FindAllSubmatch(contents, -1)
	for _, m := range allSubmatch {
		result.Items = append(result.Items, "city:"+string(m[2]))
		result.Requests = append(result.Requests, model.Request{
			Url:        string(m[1]),
			ParserFunc: ParserCity,
		})
	}
	return result
}
