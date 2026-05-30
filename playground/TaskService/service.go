package main

type TaskService struct {
	Repository TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return TaskService{Repository: r}
}

func (s *TaskService) CreateTask(title string, area Area, priority uint8, estimatedMinutes uint32) error {
	if !validateTitle(title) {
		return ErrInvalidTitle
	}
	if !validateArea(area) {
		return ErrInvalidArea
	}
	if !validatePriority(priority) {
		return ErrInvalidPriority
	}
	if !validateEstimatedMinutes(estimatedMinutes) {
		return ErrInvalidEstimatedMinutes
	}

	task := Task{
		Title:            title,
		Area:             area,
		Status:           todo,
		Priority:         priority,
		EstimatedMinutes: estimatedMinutes,
	}
	err := s.Repository.Save(&task)

	return err
}

func (s *TaskService) GetTaskById(id uint32) (Task, error) {
	return s.Repository.FindById(id)
}

func (s *TaskService) GetAllTasks() []Task {
	return s.Repository.FindAll()
}

func (s *TaskService) GetTasksByArea(area Area) ([]Task, error) {
	if validateArea(area) {
		return s.Repository.FindByArea(area), nil
	}
	return make([]Task, 0), ErrInvalidArea
}

func (s *TaskService) CompleteTask(id uint32) error {
	task, err := s.Repository.FindById(id)
	if err != nil {
		return err
	}
	if task.Status != done {
		task.Complete()
		err := s.Repository.Update(&task)
		return err
	}
	return ErrTaskAlreadyCompleted
}

func (s *TaskService) ChangePriority(id uint32, priority uint8) error {
	if !validatePriority(priority) {
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
