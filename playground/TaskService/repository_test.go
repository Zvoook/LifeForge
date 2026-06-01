package main

import "testing"

func TestRepositorySaveAssignsID(t *testing.T) {
	repository := NewRepository()

	task := Task{
		Area:             Backend,
		Title:            "Learn GO tests",
		Status:           todo,
		Priority:         8,
		EstimatedMinutes: 60,
	}

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

	task := Task{
		Area:             Backend,
		Title:            "Learn GO tests #2",
		Status:           todo,
		Priority:         7,
		EstimatedMinutes: 60,
	}

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

	firstTask := Task{
		Area:             Backend,
		Title:            "First task",
		Status:           todo,
		Priority:         5,
		EstimatedMinutes: 30,
	}

	secondTask := Task{
		Area:             English,
		Title:            "Second task",
		Status:           todo,
		Priority:         6,
		EstimatedMinutes: 40,
	}

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
	tasks := []Task{
		{
			ID:               1,
			Area:             Backend,
			Title:            "Saved task 1",
			Status:           todo,
			Priority:         5,
			EstimatedMinutes: 30,
		},
		{
			ID:               5,
			Area:             English,
			Title:            "Saved task 5",
			Status:           todo,
			Priority:         6,
			EstimatedMinutes: 45,
		},
	}

	repository := NewRepositoryFromTasks(tasks)

	newTask := Task{
		Area:             Algorithms,
		Title:            "New task",
		Status:           todo,
		Priority:         9,
		EstimatedMinutes: 90,
	}

	err := repository.Save(&newTask)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if newTask.ID != 6 {
		t.Fatalf("expected new task ID to be 6, got %d", newTask.ID)
	}
}
