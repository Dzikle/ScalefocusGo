package repository

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	_ "modernc.org/sqlite"
)

const (
	createStoryTable = "CREATE TABLE IF NOT EXISTS stories (story_id INTEGER PRIMARY KEY NOT NULL,title TEXT NOT NULL,score TEXT NOT NULL,datestamp DATETIME DEFAULT CURRENT_TIMESTAMP, UNIQUE(story_id) ON CONFLICT REPLACE);"
	insertStory      = "INSERT INTO stories (story_id,title,score) values (?,?,?)"
	selectLatesDate  = "SELECT s.datestamp FROM stories s ORDER BY s.datestamp DESC LIMIT 1"
	selectStories    = "SELECT story_id,title,score FROM stories s ORDER BY score DESC"
)

func TestGetDatestamp(t *testing.T) {

	mockDB, err := sql.Open("sqlite", ":memory:")

	if err != nil {
		t.Fatal("can't open DB")
	}

	_, err = mockDB.Exec(createStoryTable)

	if err != nil {
		t.Fatal("Cant create Table in DB")
	}

	repo := NewRepo(mockDB)
	// result := repo.GetDateStamp()
	// if result != time.Now() {
	// 	t.Fatal("cant create dateStamp")
	// }

	want := time.Now().Add(time.Hour)

	mockDB.Exec(insertStory, 0, "Test1", 15, want)
	mockDB.Exec(insertStory, 1, "test2", 10, time.Now().Add(-time.Hour))

	result := repo.GetDateStamp()

	if !result.Equal(want) {
		t.Fatalf(`Got %v, want %v.`, result, want)
	}

}

func TestDateStampSqlMock(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an err '%s' was not expected when opening a stub database connection", err)
	}
	wantedTime := time.Now().Add(time.Hour)
	defer db.Close()
	mock.ExpectQuery(selectLatesDate).WillReturnRows(sqlmock.NewRows([]string{"timevalue"}).AddRow(wantedTime))
	repo := NewRepo(db)
	result := repo.GetDateStamp()
	if !result.Equal(wantedTime) {
		t.Fatal("failed to get latest DateSTamp")
	}

}
