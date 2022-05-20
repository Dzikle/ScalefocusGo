package controller

import (
	"final/cmd/gin/entity"
	"final/cmd/gin/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListController interface {
	Save(ctx *gin.Context) entity.List
	FindAll() []entity.List
	Delete(ctx *gin.Context) error
}
type controller struct {
	service service.ListService
}

func New(service service.ListService) ListController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) entity.List {
	var list entity.List
	ctx.ShouldBindJSON(&list)
	c.service.Save(list)
	return list
}
func (c *controller) FindAll() []entity.List {
	return c.service.FindAll()
}
func (c *controller) Delete(ctx *gin.Context) error {
	var list entity.List
	list.Id, _ = strconv.Atoi(ctx.Param("id"))
	c.service.Delete(list)
	return nil
}
