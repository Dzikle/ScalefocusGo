package repository

import (
	"final/cmd/gin/entity"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository interface {
	Save(entity.List)
	Update(entity.List)
	Delete(entity.List)
	FindAll() []entity.List
	SaveTask(entity.Task)
	GetTask() []entity.Task
	UpdateTask(entity.Task)
	DeleteTask(entity.List)
}

type database struct {
	connection *gorm.DB
}

func NewRepository() Repository {
	db, err := gorm.Open(sqlite.Open("gorm.DB"), &gorm.Config{})
	if err != nil {
		panic("Failed to open DB")
	}
	db.AutoMigrate(&entity.List{}, &entity.Task{}, &entity.User{})
	database := database{connection: db}

	return &database
}

func (db *database) Save(list entity.List) {
	db.connection.Create(&list)
}
func (db *database) Update(list entity.List) {
	db.connection.Save(&list)
}
func (db *database) Delete(list entity.List) {
	db.connection.Delete(&list)
}
func (db *database) FindAll() []entity.List {
	var lists []entity.List
	db.connection.Set("gorm:auto_preload", true).Find(&lists)
	return lists
}
func (db *database) SaveTask(entity.Task) {}
func (db *database) GetTask() []entity.Task {
	return []entity.Task{}
}
func (db *database) UpdateTask(entity.Task) {}
func (db *database) DeleteTask(entity.List) {}
