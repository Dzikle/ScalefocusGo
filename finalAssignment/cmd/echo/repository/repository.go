package repository

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"final/cmd/echo/model"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"

	_ "modernc.org/sqlite"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	repo := &Repo{db: db}
	initDB(db, repo)
	return repo
}
func (r *Repo) CreateUser(username, password string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	query := "INSERT or IGNORE INTO users (username,password) values (?,?)"
	_, err = r.db.Exec(query, username, string(bytes))
	if err != nil {
		log.Println(err)
	}
}
func (r *Repo) CreateList(list model.List, name string) error {
	query := "INSERT INTO lists (name,username) values (?,?)"
	_, err := r.db.Exec(query, list.Name, name)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (r *Repo) GetLists(name string) ([]model.List, error) {
	query := fmt.Sprintf("SELECT id,name FROM lists l WHERE username = '%s' ORDER BY id DESC", name)
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	lists := []model.List{}
	for rows.Next() {
		l := model.List{}
		if err := rows.Scan(&l.Id, &l.Name); err != nil {
			log.Print(err)
		}
		lists = append(lists, l)
	}
	return lists, nil
}
func (r *Repo) DeleteList(id string) error {
	queryTask := fmt.Sprintf("DELETE FROM tasks WHERE list_id = '%s'", id)
	queryList := fmt.Sprintf("DELETE FROM lists WHERE id = %s", id)
	queryslice := []string{queryTask, queryList}

	for _, query := range queryslice {
		_, err := r.db.Exec(query)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
func (r *Repo) CreateTask(task model.Task) error {
	query := "INSERT INTO tasks (text,list_id,completed) values (?,?,?)"
	if !task.Completed {
		_, err := r.db.Exec(query, task.Text, task.ListId, 0)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	return nil
}
func (r *Repo) GetTasks(id string, username string) ([]model.Task, error) {
	query := fmt.Sprintf("SELECT t.id,t.text,t.completed FROM tasks t INNER JOIN lists l ON t.list_id = l.id WHERE list_id = %s AND l.username = '%s';", id, username)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer rows.Close()
	tasks := []model.Task{}
	for rows.Next() {
		t := model.Task{}
		if err := rows.Scan(&t.Id, &t.Text, &t.Completed); err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
func (r *Repo) PatchTask(id string, task model.Task) {
	var query string
	if task.Completed {
		query = fmt.Sprintf("UPDATE tasks SET completed = 1 WHERE id = %s", id)
	} else {
		query = fmt.Sprintf("UPDATE tasks SET completed = 0 WHERE id = %s", id)
	}
	r.db.Exec(query)
}
func (r *Repo) DeleteTask(id string) error {
	query := fmt.Sprintf("DELETE FROM tasks WHERE id = %s", id)
	_, err := r.db.Exec(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func (r *Repo) ValidateUser(username, password string) bool {
	user := model.User{}
	query := fmt.Sprintf("SELECT username,password FROM users WHERE username = '%s' ", username)
	rows, err := r.db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Fatal("Cant execute Query")
	}
	for rows.Next() {
		err = rows.Scan(&user.Username, &user.Password)
		if err != nil {
			log.Fatal("Scaning rows failed")
		}
	}
	ChckdPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.Username == username && ChckdPass == nil {
		return true
	}
	return false
}
func (r *Repo) GetWetherReq(lat, lon string) model.WeatherInfo {
	var apiKey string = "4477f4e02b1c530e57e643af6ddd3a41"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s", lat, lon, apiKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal()
	}
	weather := model.Weather{}

	json.NewDecoder(res.Body).Decode(&weather)

	result := model.WeatherInfo{
		FormatedTemp: fmt.Sprintf("%f C", weather.Main.Temp-273.15),
		Description:  weather.Weather[0].Description,
		City:         weather.City}

	return result
}
func (r *Repo) Export(user string) bool {

	lists, _ := r.GetLists(user)
	tasks := []model.Task{}
	for _, list := range lists {
		task, err := r.GetTasks(fmt.Sprintf("%d", list.Id), user)
		if err != nil {
			return false
		}
		tasks = append(tasks, task...)
	}
	csvFile, err := os.Create("Lists.csv")
	if err != nil {
		log.Printf("Failed to create file: %s", err)
		return false
	}
	writer := csv.NewWriter(csvFile)
	taskNames := []string{}
	for _, task := range tasks {
		name := task.Text
		taskNames = append(taskNames, name)
	}
	err = writer.Write(taskNames)
	if err != nil {
		log.Printf("Failed to write to the file: %s", err)
		return false
	}
	writer.Flush()
	csvFile.Close()
	return true
}
func initDB(db *sql.DB, repo *Repo) {
	queryTask := "CREATE TABLE IF NOT EXISTS lists (id INTEGER PRIMARY KEY UNIQUE,name TEXT NOT NULL, username TEXT NOT NULL);"
	queryList := "CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY,text TEXT NOT NULL,list_id INTEGER, completed INTEGER);"
	queryUser := "CREATE TABLE IF NOT EXISTS users (username TEXT PRIMARY KEY,password TEXT NOT NULL);"
	queryslice := []string{queryTask, queryList, queryUser}
	for _, query := range queryslice {
		_, err := db.Exec(query)
		if err != nil {
			log.Println(err)
		}
	}
	repo.CreateUser("tina", "222222")
	repo.CreateUser("stole", "222222")
}
