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

func validatePriority(priority uint8) bool {
	if priority < 1 || priority > 10 {
		return false
	}
	return true
}

func validateEstimatedMinutes(EstimatedMinutes uint32) bool {
	if EstimatedMinutes == 0 {
		return false
	}
	return true
}
