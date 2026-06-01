package task

import "strings"

func ValidateTitle(title string) bool {
	return strings.TrimSpace(title) != ""
}

func ValidateArea(area Area) bool {
	return area.IsValid()
}

func ValidateStatus(status Status) bool {
	return status.IsValid()
}

func ValidatePriority(priority int) bool {
	if priority < 1 || priority > 10 {
		return false
	}
	return true
}

func ValidateEstimatedMinutes(em int) bool {
	if em <= 0 {
		return false
	}
	return true
}
