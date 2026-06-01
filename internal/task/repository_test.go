package task

import "testing"

func TestRepositorySaveAssignsID(t *testing.T) {
	repository := NewRepository()
	task := NewTestTask("Learn GO tests")

	err := repository.Save(&task)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if task.ID != 1 {
		t.Fatalf("expected task ID to be 1, got %v", task.ID)
	}
}

func TestRepositoryFindByIDReturnsSavedTasks(t *testing.T) {
	repository := NewRepository()

	task := NewTestTask("Learn GO tests")

	err := repository.Save(&task)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	foundTask, err := repository.FindById(task.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if foundTask.Title != task.Title {
		t.Fatalf("expected title %q, got %q", task.Title, foundTask.Title)
	}
}

func TestRepositoryFindByIDWhenTaskDoesNotExist(t *testing.T) {
	repository := NewRepository()

	_, err := repository.FindById(999)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err != ErrTaskNotFound {
		t.Fatalf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestRepositoryFindAllReturnsTasksSortedByID(t *testing.T) {
	repository := NewRepository()

	firstTask := NewTestTask("First task")
	secondTask := NewTestTask("Second task")

	err := repository.Save(&firstTask)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = repository.Save(&secondTask)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	tasks := repository.FindAll()

	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].ID != 1 {
		t.Fatalf("expected first task ID to be 1, got %d", tasks[0].ID)
	}

	if tasks[1].ID != 2 {
		t.Fatalf("expected second task ID to be 2, got %d", tasks[1].ID)
	}
}

func TestNewRepositoryFromTasksRestoresNextID(t *testing.T) {
	first := NewTestTask("Saved task 1")
	first.ID = 1

	second := NewTestTask("Saved task 5")
	second.ID = 5

	tasks := []Task{
		first,
		second,
	}

	repository := NewRepositoryFromTasks(tasks)

	newTask := NewTestTask("New task")

	err := repository.Save(&newTask)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if newTask.ID != 6 {
		t.Fatalf("expected new task ID to be 6, got %d", newTask.ID)
	}
}
