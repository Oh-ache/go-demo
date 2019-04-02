package parser

import (
	"learngo/single/engine"
	"learngo/single/model"
	"regexp"
	"strings"
)

var nameRe = regexp.MustCompile(`<h1 class="nickName"[^>]*>([^<]*)</h1>`)
var imageRe = regexp.MustCompile(`<div class="logo f-fl" style="background-image:url[(]([^?]*)[?]scrop=1&amp;crop=1&amp;cpos=north&amp;w=200&amp;h=200[)];" data-v-5b109fc3></div>`)
var idRe = regexp.MustCompile(`<div class="id"[^>]*>ID：([^<]*)</div>`)
var infoRe = regexp.MustCompile(`<div class="des f-cl"[^>]*>([^<]*)</div>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = extractString(contents, nameRe)
	profile.Image = extractString(contents, imageRe)
	profile.Id = extractString(contents, idRe)
	info := extractString(contents, infoRe)

	if info != "" {
		arr := strings.Split(info, " | ")
		if len(arr) > 2 {
			profile.Address = string(arr[0])
			profile.Age = string(arr[1])
			profile.Education = string(arr[2])
			profile.Marriage = string(arr[3])
			profile.Height = string(arr[4])
			profile.Income = string(arr[5])
		}
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
