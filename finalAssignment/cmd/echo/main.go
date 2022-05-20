package main

import (
	"database/sql"
	"final/cmd"
	hand "final/cmd/echo/handlers"
	repo "final/cmd/echo/repository"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "store.db")
	if err != nil {
		log.Fatal("can't open DB")
	}
	repo := repo.NewRepo(db)
	router := echo.New()

	router.Use(middleware.BasicAuth(hand.Basic(repo)))
	router.GET("/api", func(ctx echo.Context) error {
		return ctx.JSON(200, "Hello, World!")
	})
	router.POST("/api/lists", hand.CreateList(repo))
	router.GET("/api/lists", hand.GetLists(repo))
	router.DELETE("/api/lists/:id", hand.DeleteList(repo))

	router.POST("/api/lists/:id/tasks", hand.CreateTask(repo))
	router.GET("/api/lists/:id/tasks", hand.GetTask(repo))
	router.PATCH("/api/tasks/:id", hand.PatchTask(repo))
	router.DELETE("/api/tasks/:id", hand.DeleteTask(repo))
	router.GET("/api/weather", hand.GetWeather(repo))
	router.GET("api/list/export", hand.ExportCSV(repo))

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
