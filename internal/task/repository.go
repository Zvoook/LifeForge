package task

import "sort"

type TaskRepository interface {
	Reset()
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
	nextID uint32
}

func NewRepository() InMemoryTaskRepository {
	return InMemoryTaskRepository{tasks: make(map[uint32]Task), nextID: 1}
}

func NewRepositoryFromTasks(tasks []Task) InMemoryTaskRepository {
	repository := InMemoryTaskRepository{
		tasks:  make(map[uint32]Task),
		nextID: 1,
	}

	var maxID uint32

	for _, task := range tasks {
		repository.tasks[task.ID] = task

		if task.ID > maxID {
			maxID = task.ID
		}
	}

	repository.nextID = maxID + 1
	return repository
}

func (r *InMemoryTaskRepository) Reset() {
	r.tasks = make(map[uint32]Task)
	r.nextID = 0
}

func (r *InMemoryTaskRepository) GetId() uint32 {
	r.nextID++
	return r.nextID - 1
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
	tasks := make([]Task, 0, len(r.tasks))

	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	return tasks
}

func (r *InMemoryTaskRepository) FindByArea(a Area) []Task {
	tasks := make([]Task, 0, len(r.tasks))

	for _, task := range r.tasks {
		if task.Area == a {
			tasks = append(tasks, task)
		}
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	return tasks
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
