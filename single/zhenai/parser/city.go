package parser

import (
	"learngo/single/engine"
	"regexp"
)

const cityRe = `<a href="([^"]*)" target="_blank">([^<]*)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		engine.Run(engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseProfile,
		})
	}
	//for _, m := range matches {
	//	result.Items = append(result.Items, string(m[2]))
	//	result.Requests = append(result.Requests,
	//		engine.Request{
	//			Url:        string(m[1]),
	//			ParserFunc: engine.NilParser,
	//		})
	//}
	return result
}
