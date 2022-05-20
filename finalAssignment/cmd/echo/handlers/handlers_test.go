package handlers

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"final/cmd/echo/model"
	repo "final/cmd/echo/repository"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func initTestDB(t *testing.T) (*repo.Repo, *sql.DB) {
	mockDB, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("can't open DB")
	}
	return repo.NewRepo(mockDB), mockDB
}

func TestGetLists(t *testing.T) {
	repo, mockDB := initTestDB(t)
	wantList := []model.List{
		{
			Id:   1,
			Name: "test"}}
	query := "INSERT INTO lists (name,username) values (?,?)"
	_, err := mockDB.Exec(query, wantList[0].Name, "stole")
	if err != nil {
		log.Println(err)
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/lists", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Request().Header.Set("username", "stole")

	handler := GetLists(repo)(c)
	gotList := []model.List{}
	json.NewDecoder(rec.Body).Decode(&gotList)

	if assert.NoError(t, handler) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, wantList, gotList)
	}
}
func TestCreateList(t *testing.T) {
	repo, _ := initTestDB(t)
	wantList := model.List{
		Id:   1,
		Name: "test"}
	jsonBytes, err := json.Marshal(wantList)
	if err != nil {
		log.Println(err)
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/lists", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.Request().Header.Set("username", "stole")

	handler := CreateList(repo)(c)

	gotList, _ := repo.GetLists(c.Request().Header.Get("username"))

	if assert.NoError(t, handler) {
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, wantList.Name, gotList[0].Name)
	}
}
func TestDeleteList(t *testing.T) {
	repo, db := initTestDB(t)
	query := "INSERT INTO lists (name,username) values (?,?)"
	user := "stole"
	_, err := db.Exec(query, "Test", user)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}
	e := echo.New()
	req := httptest.NewRequest("DELETE", "/api/lists/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Request().Header.Set("username", "stole")
	c.SetParamNames("id")
	c.SetParamValues("1")
	handler := DeleteList(repo)(c)
	assert.NoError(t, handler)
}
func TestCreateTask(t *testing.T) {
	repo, _ := initTestDB(t)
	wantTask := model.Task{
		Id:        1,
		Text:      "Test",
		ListId:    1,
		Completed: false,
	}
	jsonBytes, err := json.Marshal(wantTask)
	if err != nil {
		log.Println(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/lists/:id/tasks", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, res)
	c.Request().Header.Set("username", "stole")
	c.SetParamNames("id")
	c.SetParamValues("1")
	handler := CreateTask(repo)(c)
	if assert.NoError(t, handler) {
		assert.Equal(t, http.StatusCreated, res.Code)
	}
}
func TestGetTasks(t *testing.T) {
	repo, mockDB := initTestDB(t)
	wantList := []model.Task{
		{
			Text:      "TEST",
			ListId:    1,
			Completed: true}}

	query := "INSERT INTO tasks (text,list_id,completed) values (?,?,?)"
	_, err := mockDB.Exec(query, wantList[0].Text, wantList[0].ListId, 0)
	if err != nil {
		log.Println(err)
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/lists/:id/tasks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Request().Header.Set("username", "stole")
	c.SetParamNames("id")
	c.SetParamValues("1")
	handler := GetTask(repo)
	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
func TestDeleteTask(t *testing.T) {
	repo, db := initTestDB(t)
	query := "INSERT INTO tasks (text,list_id,completed) values (?,?,?)"
	_, err := db.Exec(query, "Test", 1, 0)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}
	e := echo.New()
	req := httptest.NewRequest("DELETE", "/api/tasks/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Request().Header.Set("username", "stole")
	c.SetParamNames("id")
	c.SetParamValues("1")
	handler := DeleteTask(repo)(c)

	assert.NoError(t, handler)
}
func TestGetWeather(t *testing.T) {
	repo, _ := initTestDB(t)
	want := model.WeatherInfo{
		City: "Vinica",
	}
	req := httptest.NewRequest(http.MethodPost, "/api/weather", nil)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, res)
	c.Request().Header.Set("lat", "41.8828")
	c.Request().Header.Set("lon", "22.5092")

	handler := GetWeather(repo)(c)
	if assert.NoError(t, handler) {
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Contains(t, res.Body.String(), want.City)
	}
}
func TestPatchTask(t *testing.T) {
	repo, _ := initTestDB(t)
	wantTask := model.Task{
		Id:        1,
		Text:      "Test",
		ListId:    1,
		Completed: false,
	}
	jsonBytes, err := json.Marshal(wantTask)
	if err != nil {
		log.Println(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/tasks/:id", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, res)
	c.Request().Header.Set("username", "tina")
	c.SetParamNames("id")
	c.SetParamValues("1")
	handler := PatchTask(repo)(c)
	if assert.NoError(t, handler) {
		assert.Equal(t, http.StatusOK, res.Code)
	}
}
func TestBasicAuth(t *testing.T) {
	repo, _ := initTestDB(t)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	h := middleware.BasicAuth(Basic(repo))(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})
	basic := "basic"
	assert := assert.New(t)
	// Valid credentials
	auth := basic + " " + base64.StdEncoding.EncodeToString([]byte("stole:222222"))
	req.Header.Set(echo.HeaderAuthorization, auth)
	assert.NoError(h(c))

	// Valid credentials
	auth = basic + " " + base64.StdEncoding.EncodeToString([]byte("stole:222222"))
	req.Header.Set(echo.HeaderAuthorization, auth)
	assert.NoError(h(c))

	// Case-insensitive header scheme
	auth = strings.ToUpper(basic) + " " + base64.StdEncoding.EncodeToString([]byte("stole:222222"))
	req.Header.Set(echo.HeaderAuthorization, auth)
	assert.NoError(h(c))

	// Invalid credentials
	auth = basic + " " + base64.StdEncoding.EncodeToString([]byte("stole:invalid-password"))
	req.Header.Set(echo.HeaderAuthorization, auth)
	he := h(c).(*echo.HTTPError)
	assert.Equal(http.StatusUnauthorized, he.Code)

	// Missing Authorization header
	req.Header.Del(echo.HeaderAuthorization)
	he = h(c).(*echo.HTTPError)
	assert.Equal(http.StatusUnauthorized, he.Code)

	// Invalid Authorization header
	auth = base64.StdEncoding.EncodeToString([]byte("invalid"))
	req.Header.Set(echo.HeaderAuthorization, auth)
	he = h(c).(*echo.HTTPError)
	assert.Equal(http.StatusUnauthorized, he.Code)
}
func TestGetUser(t *testing.T) {
	repo, _ := initTestDB(t)
	req := httptest.NewRequest(http.MethodPost, "/api/weather", nil)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, res)

	Basic(repo)("test", "testPassword", c)

	want := "test"
	got := GetUser(c)

	assert.Equal(t, want, got)
}
