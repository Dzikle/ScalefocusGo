package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

type Stories struct {
	storiesId []int
}
type Story struct {
	Title string `json:"title,omitempty"`
	Score int    `json:"score,omitempty"`
}

var wg sync.WaitGroup

func FetchTopStories() *Stories {
	u, err := url.Parse("https://hacker-news.firebaseio.com/v0/topstories.json")
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(err)
	}
	stories := Stories{}
	json.NewDecoder(res.Body).Decode(&stories.storiesId)
	return &stories
}

func (s *Stories) FetchTop10() chan Story {
	ch := make(chan Story, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			u, err := url.Parse("https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(s.storiesId[i]) + ".json")

			req, err := http.NewRequest("GET", u.String(), nil)
			if err != nil {
				log.Fatal(err)
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			if res.StatusCode != http.StatusOK {
				log.Fatal(err)
			}
			story := Story{}
			json.NewDecoder(res.Body).Decode(&story)
			ch <- story
		}(i)
	}
	return ch
}

func HandleListen() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		topStoriesIDs := FetchTopStories()

		var stories []Story

		ch := topStoriesIDs.FetchTop10()

		for i := 0; i < 10; i++ {
			story := <-ch
			stories = append(stories, story)

		}

		result, err := json.MarshalIndent(stories, "", "")
		if err != nil {
			http.Error(w, "Not a good json file", http.StatusBadRequest)
			return
		}

		w.Write([]byte(result))
	}
}

func main() {
	router := http.NewServeMux()
	router.Handle("/top", HandleListen())
	http.ListenAndServe(":8080", router)
}
