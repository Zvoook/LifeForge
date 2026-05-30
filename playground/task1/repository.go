package main

type TaskRepository interface {
	Save(task *Task) error
	FindById(id uint32) (Task, error)
	FindAll() []Task
	FindByArea(area Area) []Task
	Update(task *Task) error
	Delete(id uint32) error
}

type InMemoryTaskRepository struct{
	tasks map[uint32]Task
	nextId uint32
}

func (r* InMemoryTaskRepository) getId() uint32 {
	r.nextId++
	return r.nextId
}

func (r* InMemoryTaskRepository) Save (t *Task) error {
	id := r.getId()
	t.id = id
	r.tasks[id] = *t
	return nil
}

func (r* InMemoryTaskRepository) FindById (id uint32) (Task, error) {
	task, contains := r.tasks[id]
	if contains {
		return task, nil
	}
	return Task{}, ErrTaskNotFound
}

func (r* InMemoryTaskRepository)