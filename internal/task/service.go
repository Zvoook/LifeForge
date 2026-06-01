package task

import "strings"

type TaskService struct {
	Repository TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return TaskService{Repository: r}
}

func (s *TaskService) CreateTask(title string, area Area, priority int, estimatedMinutes int) (Task, error) {
	title = strings.TrimSpace(title)

	if !ValidateTitle(title) {
		return Task{}, ErrInvalidTitle
	}
	if !ValidateArea(area) {
		return Task{}, ErrInvalidArea
	}
	if !ValidatePriority(priority) {
		return Task{}, ErrInvalidPriority
	}
	if !ValidateEstimatedMinutes(estimatedMinutes) {
		return Task{}, ErrInvalidEstimatedMinutes
	}

	task := Task{
		Title:            title,
		Area:             area,
		Status:           Todo,
		Priority:         priority,
		EstimatedMinutes: estimatedMinutes,
	}
	err := s.Repository.Save(&task)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *TaskService) GetTaskById(id uint32) (Task, error) {
	return s.Repository.FindById(id)
}

func (s *TaskService) GetAllTasks() []Task {
	return s.Repository.FindAll()
}

func (s *TaskService) GetTasksByArea(area Area) ([]Task, error) {
	if ValidateArea(area) {
		return s.Repository.FindByArea(area), nil
	}
	return make([]Task, 0), ErrInvalidArea
}

func (s *TaskService) CompleteTask(id uint32) error {
	task, err := s.Repository.FindById(id)
	if err != nil {
		return err
	}
	if task.Status != Done {
		task.Complete()
		err := s.Repository.Update(&task)
		return err
	}
	return ErrTaskAlreadyCompleted
}

func (s *TaskService) ChangePriority(id uint32, priority int) error {
	if !ValidatePriority(priority) {
		return ErrInvalidPriority
	}
	task, err := s.Repository.FindById(id)
	if err != nil {
		return err
	}
	task.ChangePriority(priority)
	err = s.Repository.Update(&task)
	return err
}

func (s *TaskService) DeleteTask(id uint32) error {
	return s.Repository.Delete(id)
}

func (s *TaskService) ClearAll() {
	s.Repository.Reset()
}
