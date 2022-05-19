package main

import (
	"database/sql"
	repository "l22/Repository"
	. "l22/handlers"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	router := http.NewServeMux()

	db, err := sql.Open("sqlite", "store.db")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS stories (story_id INTEGER PRIMARY KEY NOT NULL,title TEXT NOT NULL,score TEXT NOT NULL,datestamp DATETIME DEFAULT CURRENT_TIMESTAMP, UNIQUE(story_id) ON CONFLICT REPLACE);")
	statement.Exec()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepo(db)

	url := "https://hacker-news.firebaseio.com/v0/topstories.json"
	s := Stories{Url: "https://hacker-news.firebaseio.com/v0/item/", Count: 10}
	router.Handle("/api/top", HandleListen(url, &s, repo))
	router.Handle("/top", HandleTopStories(url, &s, repo))

	http.ListenAndServe(":8090", router)
}
