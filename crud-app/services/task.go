package services

import (
	"crud-app/models"
	"fmt"
)

type TaskService interface {
	CreateTask(title, body string) (models.Task, error)
	GetTasks() []models.Task
	GetTaskById(id int) (models.Task, error)
	UpdateTask(id int, title, body string) (models.Task, error)
	DeleteTask(id int) error
}

type InMemoryTaskService struct {
	tasks []models.Task
	currentID int
}

func NewInMemoryTaskService() *InMemoryTaskService {
	return &InMemoryTaskService{}
}

func (s *InMemoryTaskService) CreateTask(title, body string) (models.Task, error){
	s.currentID++
	task := models.Task{
		ID: s.currentID,
		Title: title,
		Body: body,
	}
	s.tasks =append(s.tasks, task)
	return task, nil
}

func (s *InMemoryTaskService) GetTasks() []models.Task {
	return s.tasks
}

func (s *InMemoryTaskService) GetTaskById(id int) (models.Task, error){
	for _, task := range s.tasks {
		if task.ID == id{
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("task not found")
}

func (s *InMemoryTaskService) UpdateTask(id int, title, body string) (models.Task, error) {
	for i := 0; i < len(s.tasks); i++ {
		if s.tasks[i].ID == id {
			s.tasks[i].Title = title
			s.tasks[i].Body = body
			return s.tasks[i], nil
		}
	}
	return models.Task{}, fmt.Errorf("task not found")
}

func (s *InMemoryTaskService) DeleteTask(id int) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i + 1:]...)
			return nil
		}
	}
	return fmt.Errorf("task not found")
} 
