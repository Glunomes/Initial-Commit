package app

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type TaskService interface {
	Save(t domain.Task) (domain.Task, error)
	GetAll() ([]domain.Task, error)
	GetByID(id uint64) (domain.Task, error)
	Update(t domain.Task) (domain.Task, error)
	Delete(id uint64) error
}

type taskService struct {
	taskRepo database.TaskRepository
}

func NewTaskService(tr database.TaskRepository) TaskService {
	return &taskService{
			taskRepo: tr,
	}
}

func (s *taskService) Save(t domain.Task) (domain.Task, error) {
	return s.taskRepo.Save(t)
}

func (s *taskService) GetAll() ([]domain.Task, error) {
	return s.taskRepo.FindAll()
}

func (s *taskService) GetByID(id uint64) (domain.Task, error) {
	return s.taskRepo.FindById(id)
}

func (s *taskService) Update(t domain.Task) (domain.Task, error) {
	return s.taskRepo.Update(t)
}

func (s *taskService) Delete(id uint64) error {
	return s.taskRepo.Delete(id)
}
