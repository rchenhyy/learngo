package main

import (
	"fmt"
	"rchenhyy/learngo/interfaces/mock"
	"rchenhyy/learngo/interfaces/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com/")
}

func post(p Poster) {
	p.Post("https://www.imooc.com/", map[string]string{
		"name":   "rchenyy",
		"course": "golang",
	})
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}

func main() {
	var r Retriever
	r = &mock.Retriever{Contents: "Hello world"}
	fmt.Println(download(r))
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	//fmt.Println(download(r))
	inspect(r)

	// Type assertion
	realR, ok := r.(*real.Retriever)
	if ok {
		fmt.Println(realR.UserAgent)
	} else {
		fmt.Println("not a real retriever")
	}
}
