package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

type Stories struct {
	StoriesId []int
	Url       string
	Count     int
}
type Story struct {
	Id    int
	Title string
	Score int
}
type PageData struct {
	PageTitle string
	Links     []Story
}

type Storage interface {
	GetDateStamp() time.Time
	GetStories() []Story
	SaveStories(list []Story)
	Delete()
}

var wg sync.WaitGroup

//"https://hacker-news.firebaseio.com/v0/topstories.json"

func (st Stories) FetchStoriesID(s string) *Stories {

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
	json.NewDecoder(res.Body).Decode(&st.StoriesId)
	return &st
}

func (s *Stories) FetchTop10() []Story {
	ch := make(chan Story, s.Count)
	var Stories []Story
	for i := 0; i < s.Count; i++ {
		go func(i int) {
			u, err := url.Parse(s.Url + strconv.Itoa(s.StoriesId[i]) + ".json")
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
			story.Id = s.StoriesId[i]
			ch <- story
		}(i)
	}
	defer close(ch)

	for i := 0; i < s.Count; i++ {
		Stories = append(Stories, <-ch)
	}

	return Stories
}

func HandleListen(sUrl string, storiesI *Stories, storage Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}

		var stories []Story

		beforeHour := time.Now().Add(-time.Hour)
		if storage.GetDateStamp().Before(beforeHour) {
			stories = storiesI.FetchStoriesID(sUrl).FetchTop10()
			storage.SaveStories(stories)

		} else {
			stories = storage.GetStories()
		}

		result, err := json.MarshalIndent(stories, "", "")
		if err != nil {
			http.Error(w, "Not a good json file", http.StatusBadRequest)
			return
		}

		w.Write(result)
	}
}

func HandleTopStories(s string, storiesI *Stories, storage Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}

		var stories []Story

		beforeHour := time.Now().Add(-time.Hour)
		nowHour := storage.GetDateStamp()
		if nowHour.Before(beforeHour) {
			stories = storiesI.FetchStoriesID(s).FetchTop10()
			storage.SaveStories(stories)

		} else {
			stories = storage.GetStories()
		}

		tmpl := template.Must(template.ParseFiles("topstories.html"))
		data := PageData{
			PageTitle: "Top 10 Stories from HackerNews",
			Links:     stories,
		}
		tmpl.Execute(w, data)
	}
}
