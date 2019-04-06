package main

import (
	"fmt"
	"go-demo/retriever/mock"
	"go-demo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func post(poster Poster) {
	poster.Post("https://www.imooc.com",
		map[string]string{
			"name":   "ache",
			"course": "golang",
		})
}

const url = "https://www.imooc.com"

type RetriverPoster interface {
	Retriever
	Poster
}

func session(s RetriverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked in imooc",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspect", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Printf("> Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever
	r = &mock.Retriever{"this is a fake imooc"}
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Moilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	//fmt.Println(download(r))

	if realTriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realTriever.UserAgent)
	} else {
		fmt.Println("is not real triever")
	}

	retriever := mock.Retriever{"this is a fake imooc"}
	fmt.Println("start session")
	fmt.Println(session(&retriever))
}
