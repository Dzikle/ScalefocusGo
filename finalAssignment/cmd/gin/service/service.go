package service

import (
	"final/cmd/gin/entity"
	"final/cmd/gin/repository"
)

type ListService interface {
	Save(entity.List) entity.List
	FindAll() []entity.List
	Update(entity.List)
	Delete(entity.List)
}

type listService struct {
	repository repository.Repository
}

func New(repo repository.Repository) ListService {
	return &listService{
		repository: repo,
	}
}
func (service *listService) Save(list entity.List) entity.List {
	service.repository.Save(list)
	return list
}
func (service *listService) FindAll() []entity.List {
	return service.repository.FindAll()
}
func (service *listService) Delete(list entity.List) {
	service.repository.Delete(list)
}
func (service *listService) Update(list entity.List) {
	service.repository.Update(list)
}
