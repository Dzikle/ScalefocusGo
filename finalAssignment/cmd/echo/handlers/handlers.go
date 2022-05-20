package handlers

import (
	"final/cmd/echo/model"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

type Repo interface {
	CreateList(model.List, string) error
	GetLists(string) ([]model.List, error)
	DeleteList(string) error
	CreateTask(model.Task) error
	GetTasks(string, string) ([]model.Task, error)
	PatchTask(string, model.Task)
	DeleteTask(string) error
	ValidateUser(string, string) bool
	Export(string) bool
	GetWetherReq(lat, lon string) model.WeatherInfo
}

func GetUser(c echo.Context) string {
	return c.Request().Header.Get("username")
}
func CreateList(r Repo) echo.HandlerFunc {
	return func(c echo.Context) error {
		list := model.List{}
		err := c.Bind(&list)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		err = r.CreateList(list, GetUser(c))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusCreated, list)
	}
}
func GetLists(r Repo) echo.HandlerFunc {
	return func(c echo.Context) error {
		lists, err := r.GetLists(GetUser(c))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, lists)
	}
}
func DeleteList(r Repo) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := r.DeleteList(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, "The lists is deleted")
	}
}
func CreateTask(r Repo) echo.HandlerFunc {
	return func(c echo.Context) error {
		task := model.Task{}
		id := c.Param("id")
		task.ListId, _ = strconv.Atoi(id)
		err := c.Bind(&task)
		if err != nil && task.Text != "" {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		err = r.CreateTask(task)
		if err != nil && task.Text != "" {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusCreated, task)
	}
}
func GetTask(r Repo) echo.HandlerFunc {
	return func(c echo.Context) error {
		tasks, err := r.GetTasks(c.Param("id"), GetUser(c))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, tasks)
	}
}
func PatchTask(r Repo) echo.HandlerFunc {
	return func(c echo.Context) error {
		task := model.Task{}
		err := c.Bind(&task)
		if err != nil {
			log.Fatal(err)
		}
		r.PatchTask(c.Param("id"), task)
		return c.JSON(http.StatusOK, &task)
	}
}
func DeleteTask(r Repo) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := r.DeleteTask(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, "The task is deleted")
	}
}
func Basic(r Repo) middleware.BasicAuthValidator {
	return func(username, password string, ctx echo.Context) (bool, error) {
		ctx.Request().Header.Set("username", username)
		if r.ValidateUser(username, password) {
			return true, nil
		}
		return false, echo.ErrUnauthorized
	}
}
func GetWeather(r Repo) echo.HandlerFunc {
	return func(c echo.Context) error {
		WeatherInfo := r.GetWetherReq(c.Request().Header.Get("lat"), c.Request().Header.Get("lon"))
		if WeatherInfo.Description != "" && WeatherInfo.City != "" {
			return c.JSON(http.StatusOK, WeatherInfo)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
}
func ExportCSV(r Repo) echo.HandlerFunc {
	return func(c echo.Context) error {
		if r.Export(GetUser(c)) {
			return c.JSON(http.StatusOK, "successful operation")
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
}
