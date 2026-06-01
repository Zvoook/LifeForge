package task

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveAndLoadTasksFromFile(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "save.json")

	tasks := []Task{
		NewTestTask("task_1"),
		NewTestTask("task_2"),
	}

	err := SaveTasksToFile(tasks, filename)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	loaded, err := LoadTasksFromFile(filename)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(loaded) != len(tasks) {
		t.Fatalf("expected %v tasks, got %v", len(tasks), len(loaded))
	}
	if loaded[0].ID != tasks[0].ID {
		t.Fatalf("expected first task ID %d, got %d", tasks[0].ID, loaded[0].ID)
	}
	if loaded[0].Title != tasks[0].Title {
		t.Fatalf("expected first task title %q, got %q", tasks[0].Title, loaded[0].Title)
	}
	if loaded[1].Status != tasks[1].Status {
		t.Fatalf("expected second task status %v, got %v", tasks[1].Status, loaded[1].Status)
	}
}

func TestLoadFromFileReturnsEmptySliceWhenFileDoesNotExist(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "non-save.json")

	loaded, err := LoadTasksFromFile(filename)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(loaded) != 0 {
		t.Fatalf("expected 0 tasks, got %v", len(loaded))
	}
}

func TestLoadTasksFromFileReturnsErrorForInvalidJSON(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "broken-save.json")

	err := os.WriteFile(filename, []byte("{ invalid json"), 0644)
	if err != nil {
		t.Fatalf("expected no error while preparing test file, got %v", err)
	}

	_, err = LoadTasksFromFile(filename)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
