package main

type TaskRepository interface {
	GetId() uint32
	Save(task *Task) error
	FindById(id uint32) (Task, error)
	FindAll() []Task
	FindByArea(area Area) []Task
	Update(task *Task) error
	Delete(id uint32) error
}

type InMemoryTaskRepository struct {
	tasks  map[uint32]Task
	nextId uint32
}

func NewRepository() InMemoryTaskRepository {
	return InMemoryTaskRepository{tasks: make(map[uint32]Task), nextId: 0}
}

func (r *InMemoryTaskRepository) GetId() uint32 {
	r.nextId++
	return r.nextId
}

func (r *InMemoryTaskRepository) Save(t *Task) error {
	id := r.GetId()
	t.ID = id
	r.tasks[id] = *t
	return nil
}

func (r *InMemoryTaskRepository) FindById(id uint32) (Task, error) {
	task, contains := r.tasks[id]

	if contains {
		return task, nil
	}

	return Task{}, ErrTaskNotFound
}

func (r *InMemoryTaskRepository) FindAll() []Task {
	founded := make([]Task, 0, len(r.tasks))

	for _, task := range r.tasks {
		founded = append(founded, task)
	}

	return founded
}

func (r *InMemoryTaskRepository) FindByArea(a Area) []Task {
	founded := make([]Task, 0, len(r.tasks))

	for _, task := range r.tasks {
		if task.Area == a {
			founded = append(founded, task)
		}
	}

	return founded
}

func (r *InMemoryTaskRepository) Update(t *Task) error {
	id := t.ID
	_, contains := r.tasks[id]

	if contains {
		r.tasks[id] = *t
		return nil
	}

	return ErrTaskNotFound
}

func (r *InMemoryTaskRepository) Delete(id uint32) error {
	_, contains := r.tasks[id]
	if !contains {
		return ErrTaskNotFound
	}

	delete(r.tasks, id)
	return nil
}
