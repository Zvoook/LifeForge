package task

type Area int

const (
	Backend Area = iota
	English
	Guitar
	Algorithms
	University
	Unknown
)

func (a Area) String() string {
	switch a {
	case Backend:
		return "Backend"
	case English:
		return "English"
	case Guitar:
		return "Guitar"
	case Algorithms:
		return "Algorithms"
	case University:
		return "University"
	default:
		return "Unknown Area"
	}
}

func (a *Area) IsValid() bool {
	switch *a {
	case Backend, English, Guitar, Algorithms, University:
		return true
	default:
		return false
	}
}
