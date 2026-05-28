package main

type TaskRepository interface {
	Save(task *Task) error
	FindById(id uint32) (Task, error)
	FindAll() []Task
	FindByArea(area Area) []Task
	Update(task *Task) error
	Delete(id uint32) error
}

type InMemoryTaskRepository struct {
	tasks map[uint32]Task
}
