package parser

import (
	"regexp"

	"github.com/GoSpider/crawler/engine"
)

// match[1]=url match[2]=name
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParserCity(contents []byte) engine.ParserResult {
	result := engine.ParserResult{}
	re := regexp.MustCompile(cityRe)
	allSubmatch := re.FindAllSubmatch(contents, -1)
	for _, m := range allSubmatch {
		name := string(m[2])
		result.Items = append(result.Items, "user "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParseProfile(c, name) // 函数式编程，使用函数包裹函数
			},
		})
	}
	return result
}
