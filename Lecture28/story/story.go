package story

import (
	"Lecture28/db"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type StoryService struct {
	urlBase string
	repo    Repository
}
type Story struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Score int    `json:"score"`
}

type Repository interface {
	SaveStories(ctx context.Context, data db.SaveStoriesParams) (sql.Result, error)
}

func NewStoryService(url string, repo Repository) *StoryService {
	return &StoryService{urlBase: url, repo: repo}
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
		res1 := db.SaveStoriesParams{Storyid: int32(v.Id), Title: v.Title, Score: int32(v.Score), Timestamp: time.Now()}
		ss.repo.SaveStories(context.Background(), res1)
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
