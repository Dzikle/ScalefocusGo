package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Stories struct {
	storiesId []int
	url       string
}
type Story struct {
	id    int
	Title string
	Score int
}
type PageData struct {
	PageTitle string
	Links     []Story
}

//"https://hacker-news.firebaseio.com/v0/topstories.json"

func (st *Stories) FetchTopStories(s string) *Stories {

	u, err := url.Parse(s)
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
	json.NewDecoder(res.Body).Decode(&st.storiesId)
	return st
}

func (s *Stories) FetchTop10() chan Story {
	ch := make(chan Story, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			u, err := url.Parse(s.url + strconv.Itoa(s.storiesId[i]) + ".json")
			// u := s.url + "/" + strconv.Itoa(s.storiesId[i]) + ".json"
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
			story.id = s.storiesId[i]
			ch <- story
		}(i)
	}

	return ch
}

func HandleListen(sUrl string, storiesI *Stories) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		storiesI.FetchTopStories(sUrl)
		var stories []Story
		ch := storiesI.FetchTop10()

		for i := 0; i < 10; i++ {
			story := <-ch
			stories = append(stories, story)

		}

		result, err := json.MarshalIndent(stories, "", "")
		if err != nil {
			http.Error(w, "Not a good json file", http.StatusBadRequest)
			return
		}

		w.Write(result)
	}
}

func HandleTopStories(s string, storiesI *Stories) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		storiesI.FetchTopStories(s)

		tmpl := template.Must(template.ParseFiles("topstories.html"))

		var Links []Story

		ch := storiesI.FetchTop10()

		for i := 0; i < 10; i++ {
			story := <-ch
			Links = append(Links, story)
		}
		data := PageData{
			PageTitle: "Top 10 Stories from HackerNews",
			Links:     Links,
		}
		tmpl.Execute(w, data)
	}
}

func main() {
	router := http.NewServeMux()
	url := "https://hacker-news.firebaseio.com/v0/topstories.json"
	s := Stories{url: "https://hacker-news.firebaseio.com/v0/item/"}
	router.Handle("/api/top", HandleListen(url, &s))
	router.Handle("/top", HandleTopStories(url, &s))

	http.ListenAndServe(":8080", router)
}
