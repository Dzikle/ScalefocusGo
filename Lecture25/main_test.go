package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

func mockGetStories(ids []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(ids)
	}
}

func TestHandleTop10(t *testing.T) {

	s1 := Story{Title: "Test1", Score: 122, id: 12}
	s2 := Story{Title: "Test2", Score: 423, id: 24}
	s3 := Story{Title: "Test3", Score: 15, id: 242}
	s4 := Story{Title: "Test4", Score: 55, id: 24332}
	s5 := Story{Title: "Test5", Score: 27, id: 2452}
	s6 := Story{Title: "Test6", Score: 978, id: 2462}
	s7 := Story{Title: "Test7", Score: 23, id: 245642}
	s8 := Story{Title: "Test8", Score: 111, id: 2442}
	s9 := Story{Title: "Test9", Score: 222, id: 24782}
	s10 := Story{Title: "Test10", Score: 441, id: 24122}
	stories := []Story{s1, s2, s3, s4, s5, s6, s7, s8, s9, s10}
	var sList Stories
	for _, story := range stories {
		sList.storiesId = append(sList.storiesId, story.id)
	}

	router := http.NewServeMux()
	mockServer := httptest.NewServer(router)
	router.Handle("/api/top", mockGetStories(sList.storiesId))
	router.Handle("/", mockPostTop10(stories, mockServer.URL))
	sList.url = mockServer.URL + "/"
	res := sList.FetchTopStories(mockServer.URL + "/api/top")
	if reflect.DeepEqual(res, sList) {
		t.Fatalf(`Got %v, want %v.`, res, sList)
	}
	chRes := sList.FetchTop10()
	var res2 []Story

	for i := 0; i < 10; i++ {
		story := <-chRes
		res2 = append(res2, story)
	}
	if reflect.DeepEqual(res2, stories) {
		t.Fatalf(`Got %v, want %v.`, res2, stories)
	}

}
func mockPostTop10(stories []Story, url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, story := range stories {
			if r.URL.Path == "/"+strconv.Itoa(story.id)+".json" {
				json.NewEncoder(w).Encode(story)
				break
			}
		}
	}

}
