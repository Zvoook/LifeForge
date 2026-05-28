package main

type Task struct {
	id               uint32
	area             Area
	title            string
	status           Status
	priority         uint8
	estimatedMinutes uint32
}

func (t *Task) Id() uint32 {
	return t.id
}

func (t *Task) SetId(id uint32) {
	t.id = id
}

func (t *Task) Area() Area {
	return t.area
}

func (t *Task) SetArea(area Area) {
	t.area = area
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) SetTitle(title string) {
	t.title = title
}

func (t *Task) Status() Status {
	return t.status
}

func (t *Task) SetStatus(status Status) {
	t.status = status
}

func (t *Task) Priority() uint8 {
	return t.priority
}

func (t *Task) SetPriority(priority uint8) {
	t.priority = priority
}

func (t *Task) EstimatedMinutes() uint32 {
	return t.estimatedMinutes
}

func (t *Task) SetEstimatedMinutes(estimatedMinutes uint32) {
	t.estimatedMinutes = estimatedMinutes
}
