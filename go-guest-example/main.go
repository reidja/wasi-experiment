package main

import (
	"encoding/json"

	"github.com/reidja/wasi_demo/guest/cache"
	"github.com/reidja/wasi_demo/guest/console"
	"github.com/reidja/wasi_demo/guest/log"
	"github.com/reidja/wasi_demo/guest/request"
	km "github.com/reidja/wasi_demo/shared/request"
)

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	console.Println("WASM Demo App")
	console.Println(cache.Get("keyname"))
	for i := 0; i < 5; i++ {
		req := &km.Request{Url: "https://jsonplaceholder.typicode.com/posts/1"}
		resp := request.Get(req)
		post := &Post{}
		if err := json.Unmarshal([]byte(resp.Body), post); err != nil {
			log.Println("error during unmarshall of response body: " + err.Error())
		} else {
			log.Println(post.Title)
		}
	}
}
