package parser

import (
	"regexp"

	"github.com/GoSpider/chanCrawler/model"
)

// match[1]=url match[2]=name
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParserCity(contents []byte) model.ParserResult {
	result := model.ParserResult{}
	re := regexp.MustCompile(cityRe)
	allSubmatch := re.FindAllSubmatch(contents, -1)
	for _, m := range allSubmatch {
		name := string(m[2])
		result.Items = append(result.Items, "user "+name)
		result.Requests = append(result.Requests, model.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) model.ParserResult {
				return ParseProfile(c, name) // 函数式编程，使用函数包裹函数
			},
		})
	}
	return result
}
