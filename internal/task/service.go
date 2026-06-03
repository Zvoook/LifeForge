package task

import (
	"sort"
	"strings"
)

const (
	BonusPriorityMultiplier               int = 10
	BonusStatusInProgress                 int = 15
	BonusEstimatedMinutesLessThanHalfHour int = 10
	BonusEstimatedMinutesLessThanHour     int = 5
)

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

	duplicateCandidates, _ := s.GetTasksByArea(area)
	for _, candidate := range duplicateCandidates {
		if IsDuplicates(task, candidate) {
			return Task{}, ErrTaskAlreadyExist
		}
	}

	err := s.Repository.Save(&task)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *TaskService) GetTaskById(id uint32) (Task, error) {
	return s.Repository.FindByID(id)
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

func (s *TaskService) GetTasksByStatus(status Status) ([]Task, error) {
	if !ValidateStatus(status) {
		return nil, ErrInvalidStatus
	}

	tasks := s.Repository.FindAll()
	result := make([]Task, 0)

	for _, task := range tasks {
		if task.Status == status {
			result = append(result, task)
		}
	}

	return result, nil
}

func (s *TaskService) CompleteTask(id uint32) error {
	task, err := s.Repository.FindByID(id)
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

func (s *TaskService) EditPriority(id uint32, priority int) error {
	if !ValidatePriority(priority) {
		return ErrInvalidPriority
	}
	task, err := s.Repository.FindByID(id)
	if err != nil {
		return err
	}
	task.EditPriority(priority)
	err = s.Repository.Update(&task)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) EditTitle(id uint32, title string) error {
	if !ValidateTitle(title) {
		return ErrInvalidTitle
	}
	task, err := s.Repository.FindByID(id)
	if err != nil {
		return err
	}
	task.EditTitle(title)
	err = s.Repository.Update(&task)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) EditArea(id uint32, area Area) error {
	if !ValidateArea(area) {
		return ErrInvalidArea
	}
	task, err := s.Repository.FindByID(id)
	if err != nil {
		return err
	}
	task.EditArea(area)
	err = s.Repository.Update(&task)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) EditStatus(id uint32, status Status) error {
	if !ValidateStatus(status) {
		return ErrInvalidStatus
	}
	task, err := s.Repository.FindByID(id)
	if err != nil {
		return err
	}
	task.EditStatus(status)
	err = s.Repository.Update(&task)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) EditEstimatedMinutes(id uint32, minutes int) error {
	if !ValidateEstimatedMinutes(minutes) {
		return ErrInvalidEstimatedMinutes
	}
	task, err := s.Repository.FindByID(id)
	if err != nil {
		return err
	}
	task.EditEstimatedMinutes(minutes)
	err = s.Repository.Update(&task)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) DeleteTask(id uint32) error {
	return s.Repository.Delete(id)
}

func (s *TaskService) ClearAll() {
	s.Repository.Reset()
}

func (s *TaskService) EditTaskTitle(id uint32, title string) error {
	task, err := s.Repository.FindByID(id)
	if err != nil {
		return err
	}

	task.EditTitle(title)
	err = s.Repository.Update(&task)
	if err != nil {
		return err
	}
	return nil
}

type scoredTask struct {
	task  Task
	score int
}

func calculateScore(task Task) int {
	score := task.Priority * BonusPriorityMultiplier

	if task.Status == In_progress {
		score += BonusStatusInProgress
	}

	if task.EstimatedMinutes <= 30 {
		score += BonusEstimatedMinutesLessThanHalfHour
	} else if task.EstimatedMinutes <= 60 {
		score += BonusEstimatedMinutesLessThanHour
	}

	return score
}

func (s *TaskService) GetCandidates(timeLimit int) ([]scoredTask, error) {
	allTasks := s.Repository.FindAll()

	if len(allTasks) == 0 {
		return nil, ErrEmptyFolder
	}

	candidates := make([]scoredTask, 0)

	for _, candidate := range allTasks {
		if candidate.EstimatedMinutes < timeLimit && (candidate.Status == Todo || candidate.Status == In_progress) {
			candidates = append(candidates,
				scoredTask{task: candidate, score: calculateScore(candidate)})
		}
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].score > candidates[j].score || (candidates[i].score == candidates[j].score && candidates[i].task.ID < candidates[j].task.ID)
	})

	return candidates, nil
}

func (s *TaskService) BuildDailyPlan(timeLimit int) ([]Task, error, int) {
	if !ValidateEstimatedMinutes(timeLimit) {
		return nil, ErrInvalidEstimatedMinutes, 0
	}

	candidates, err := s.GetCandidates(timeLimit)
	if err != nil {
		return nil, err, 0
	}

	totalTime := 0
	result := make([]Task, 0)

	for _, candidate := range candidates {
		if totalTime+candidate.task.EstimatedMinutes <= timeLimit {
			result = append(result, candidate.task)
			totalTime += candidate.task.EstimatedMinutes
		}
	}

	return result, nil, totalTime
}
