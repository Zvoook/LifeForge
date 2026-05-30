package main

type Task struct {
	ID               uint32
	Area             Area
	Title            string
	Status           Status
	Priority         uint8
	EstimatedMinutes uint32
}
