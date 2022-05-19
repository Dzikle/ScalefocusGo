package main

import (
	hackernews "L30/hackernews"
	"net/http"
)

func main() {
	ss := hackernews.NewStoryService("https://hacker-news.firebaseio.com")
	//show the results on localhost:9000/top
	mux := http.NewServeMux()
	mux.Handle("/api/top", hackernews.HandleTopStories(*ss))
	http.ListenAndServe(":9000", mux)
}
