package hackernews

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type StoryService struct {
	urlBase string
}
type Story struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Score int    `json:"score"`
}

func NewStoryService(url string) *StoryService {
	return &StoryService{urlBase: url}
}

func (ss *StoryService) GetTopStoriesIds(maxCount int) []int {
	req, err := http.NewRequest("GET", ss.urlBase+"/v0/topstories.json", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var ids []int
	json.NewDecoder(resp.Body).Decode(&ids)
	return ids[:maxCount]
}

func (ss *StoryService) GetStories(maxCount int) []Story {
	ids := ss.GetTopStoriesIds(maxCount)
	dataChanel := make(chan Story, len(ids))
	wg := sync.WaitGroup{}
	for _, id := range ids {
		wg.Add(1)
		go func(id int) {
			dataChanel <- ss.GetStoryById(id)
			defer wg.Done()
		}(id)
	}
	wg.Wait()
	close(dataChanel)
	result := make([]Story, 0, len(ids))
	for v := range dataChanel {
		result = append(result, v)
	}
	return result
}
func (ss *StoryService) GetStoryById(id int) Story {
	url := fmt.Sprintf("%s/v0/item/%d.json", ss.urlBase, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var st Story
	json.NewDecoder(resp.Body).Decode(&st)
	return st
}

func HandleTopStories(s StoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		stList := s.GetStories(10)
		json.NewEncoder(w).Encode(stList)
	}
}
