package main

import (
	"Lecture28/db"
	handlers "Lecture28/handlers"
	"Lecture28/story"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	router := http.NewServeMux()
	dbSQlite, err := sql.Open("sqlite", "dbHN.db")
	if err != nil {
		log.Fatal(err)
	}
	query := db.New(dbSQlite)
	rez, _ := query.GetLastStoryTimeStamp(context.Background())
	fmt.Println(rez)
	apiL := handlers.API{
		StorageService: query,
		StoryService:   *story.NewStoryService("https://hacker-news.firebaseio.com", query),
	}
	router.Handle("/api/top", apiL.HandleTopStories())
	log.Println("started")
	http.ListenAndServe(":8080", router)
}
