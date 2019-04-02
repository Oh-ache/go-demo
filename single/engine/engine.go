package engine

import (
	"learngo/single/feacher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body, err := feacher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fecher: error "+
				"fetching url is %s: %v\n", r.Url, err)
			continue
		}

		log.Printf("get item %s\n", r.Url)

		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("get item %v\n", item)
		}
	}
}
