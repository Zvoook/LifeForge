package main

func validateTaskTitle(t *Task) bool {
	if t.Title() == "" {
		return false
	}
	return true
}

func validateArea(t *Task) bool {
	a := t.Area()
	return (&a).IsValid()
}

func validateTaskStatus(t *Task) bool {
	s := t.Status()
	return (&s).IsValid()
}

func validatePriority(t *Task) bool {
	if t.Priority() < 1 || t.Priority() > 10 {
		return false
	}
	return true
}

func validateEstimatedMinutes(t *Task) bool {
	if t.EstimatedMinutes() == 0 {
		return false
	}
	return true
}
