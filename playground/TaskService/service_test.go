package main

import "testing"

func TestCanNotCreateTaskWithEmptyTitle(t *testing.T) {
	service := NewTestService()

	task, err := service.CreateTask("", Backend, 10, 60)
	if err == nil {
		t.Fatalf("expected ErrInvalidTitle, got nil")
	}
	if err != ErrInvalidTitle {
		t.Fatalf("expected ErrInvalidTitle, got %v", err)
	}
	if task.Title != "" {
		t.Fatalf("expected empty task title, got %v", task.Title)
	}
}

func TestCanNotCreateTaskWithInvalidArea(t *testing.T) {
	service := NewTestService()

	_, err := service.CreateTask("task with invalid area",
		15 /*number 15 out of range in areas' list*/, 10, 60)
	if err == nil {
		t.Fatalf("expected ErrInvalidArea, got nil")
	}
	if err != ErrInvalidArea {
		t.Fatalf("expected ErrInvalidArea, got %v", err)
	}
}

func TestCanNotCompleteTaskTwice(t *testing.T) {
	service := NewTestService()

	task, err := service.CreateTask("Example task", Backend, 10, 60)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = service.CompleteTask(task.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = service.CompleteTask(task.ID)
	if err == nil {
		t.Fatalf("expected ErrTaskAlreadyCompleted, got nil")
	}
	if err != ErrTaskAlreadyCompleted {
		t.Fatalf("expected ErrTaskAlreadyCompleted, got %v", err)
	}
}

func TestCanNotCreateTaskWithInvalidPriority(t *testing.T) {
	service := NewTestService()

	_, err := service.CreateTask("task with invalid area",
		Backend, 12 /*priority out of range 1, ..., 10*/, 60)
	if err == nil {
		t.Fatalf("expected ErrInvalidPriority, got nil")
	}
	if err != ErrInvalidPriority {
		t.Fatalf("expected ErrInvalidPriority, got %v", err)
	}
}

func TestCompleteTaskChangeTaskStatus(t *testing.T) {
	service := NewTestService()

	task, err := service.CreateTask("Example task", Backend, 10, 60)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	foundedTask, _ := service.GetTaskById(task.ID)
	statusBeforeComplete := foundedTask.Status

	err = service.CompleteTask(task.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	foundedTask, _ = service.GetTaskById(task.ID)
	statusAfterComplete := foundedTask.Status
	if statusAfterComplete != done {
		t.Fatalf("expected status done after completing, got %v", statusAfterComplete)
	}
	if statusBeforeComplete == statusAfterComplete {
		t.Fatalf("expected changing status after completing, got the same status")
	}
}

func TestDeleteTaskWorkRight(t *testing.T) {
	service := NewTestService()

	task, err := service.CreateTask("Example task", Backend, 10, 60)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = service.DeleteTask(task.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = service.GetTaskById(task.ID)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err != ErrTaskNotFound {
		t.Fatalf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestTaskServiceCreateTaskTrimsTitleSpaces(t *testing.T) {
	repository := NewRepository()
	service := NewTaskService(&repository)

	task, err := service.CreateTask("   Learn Go   ", Backend, 8, 60)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if task.Title != "Learn Go" {
		t.Fatalf("expected title %q, got %q", "Learn Go", task.Title)
	}
}
