package main

func newTestService() TaskService {
	repository := NewRepository()
	return NewTaskService(&repository)
}

func newTestTask(title string) Task {
	return Task{
		Area:             Backend,
		Title:            title,
		Status:           todo,
		Priority:         5,
		EstimatedMinutes: 30,
	}
}
