package main

type Status int

const (
	todo Status = iota
	in_progress
	done
	blocked
	cancelled
)

func (s Status) String() string {
	switch s {
	case todo:
		return "To Do"
	case in_progress:
		return "In Progress"
	case done:
		return "Done"
	case blocked:
		return "Blocked"
	case cancelled:
		return "Cancelled"
	default:
		return "Unknown Status"
	}
}

func (s *Status) IsValid() bool {
	switch *s {
	case todo, in_progress, done, blocked, cancelled:
		return true
	default:
		return false
	}
}
