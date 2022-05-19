package repository

import (
	"database/sql"
	. "l22/handlers"
	"log"
	"time"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetDateStamp() time.Time {

	query := "SELECT s.datestamp FROM stories s ORDER BY s.datestamp DESC LIMIT 1"
	var datestamp time.Time
	r.db.QueryRow(query).Scan(&datestamp)
	return datestamp
}

func (r *Repo) GetStories() []Story {
	query := "SELECT story_id,title,score FROM stories s ORDER BY score DESC"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	stories := []Story{}
	for rows.Next() {
		s := Story{}
		if err := rows.Scan(&s.Id, &s.Title, &s.Score); err != nil {
			log.Print(err)
		}
		stories = append(stories, s)
	}
	return stories
}

func (r *Repo) SaveStories(list []Story) {
	query := "INSERT INTO stories (story_id,title,score) values (?,?,?)"

	for _, v := range list {
		r.db.Exec(query, v.Id, v.Title, v.Score)
	}
}

func (r *Repo) Delete() {
	query := "DELETE FROM stories"

	rows, err := r.db.Query(query)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
}
