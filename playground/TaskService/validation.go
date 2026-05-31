package main

func validateTitle(title string) bool {
	if title == "" {
		return false
	}
	return true
}

func validateArea(area Area) bool {
	return area.IsValid()
}

func validateStatus(status Status) bool {
	return status.IsValid()
}

func validatePriority(priority int) bool {
	if priority < 1 || priority > 10 {
		return false
	}
	return true
}

func validateEstimatedMinutes(em int) bool {
	if em <= 0 {
		return false
	}
	return true
}
