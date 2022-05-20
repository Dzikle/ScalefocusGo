package main

import (
	"final/cmd"
	"final/cmd/echo/middlewares"
	"final/cmd/gin/controller"
	"final/cmd/gin/repository"
	"final/cmd/gin/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	listService    service.ListService       = service.New(repository.NewRepository())
	listController controller.ListController = controller.New(listService)
)

func main() {

	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		// This is a sample demonstration of how to attach middlewares in Gin
		log.Println("Gin middleware was called")
		ctx.Next()
	})
	router.Use(middlewares.BasicAuth())

	// Add your handler (API endpoint) registrations here
	router.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello, World!")
	})
	router.GET("/api/lists", func(ctx *gin.Context) {
		ctx.JSON(200, listController.FindAll())
	})
	router.POST("/api/lists", func(ctx *gin.Context) {
		ctx.JSON(200, listController.Save(ctx))
	})
	router.DELETE("/api/lists/:id", func(ctx *gin.Context) {
		ctx.JSON(200, listController.Delete(ctx))
	})

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
