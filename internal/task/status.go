package task

type Status int

const (
	Todo Status = iota
	In_progress
	Done
	Blocked
	Cancelled
	UnknownStatus
)

func (s Status) String() string {
	switch s {
	case Todo:
		return "To Do"
	case In_progress:
		return "In Progress"
	case Done:
		return "Done"
	case Blocked:
		return "Blocked"
	case Cancelled:
		return "Cancelled"
	default:
		return "Unknown Status"
	}
}

func (s *Status) IsValid() bool {
	switch *s {
	case Todo, In_progress, Done, Blocked, Cancelled:
		return true
	default:
		return false
	}
}
