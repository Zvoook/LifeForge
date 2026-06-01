package main

func NewTestService() TaskService {
	repository := NewRepository()
	return NewTaskService(&repository)
}

func NewTestTask(title string) Task {
	return Task{
		Area:             Backend,
		Title:            title,
		Status:           todo,
		Priority:         5,
		EstimatedMinutes: 30,
	}
}
