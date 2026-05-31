package main

import (
	"fmt"
	"strings"
)

type Point2D struct {
	X    int32
	Y    int32
	GetX func() int32
	GetY func() int32
}

type Point3D struct {
	Point
	Z int32
}

type Point interface {
	GetX() int32
	GetY() int32
}

type Line struct {
	Start Point
	End   Point
	len   int32
}

type Circle struct {
	Center Point
	Radius int32
}

func (l *Line) Length() int32 {
	return l.len
}

// func (l *Line) isPositiveDirection() int8 {
// 	if l.Start.X < l.End.X {
// 		return 1 // positive x direction
// 	} else if l.Start.X > l.End.X {
// 		return -1 // negative x direction
// 	} else {
// 		return 0 // no x movement
// 	}
// }

func main() {
	// var slice []int = []int{4,5,6}
	// fmt.Printf("The slice is: %v with capacity %d ", slice, cap(slice))

	// slice = append(slice, 7, 8, 9)
	// fmt.Printf("The updated slice is: %v with capacity %d", slice, cap(slice))

	// var slice2 []int32 = make([]int32, 5, 10)
	// fmt.Printf("The second slice has length %d with capacity %d", len(slice2), cap(slice2))

	var myMap map[string]uint8 = map[string]uint8{
		"monday":    1,
		"tuesday":   2,
		"wednesday": 3,
	}

	//fmt.Printf("The map is: %v", myMap)
	//fmt.Printf("The value of key wednesday is: %d", myMap["wednesday"])

	for key, value := range myMap {
		fmt.Printf("Day: %s, Number: %d\n", key, value)
	}

	//for i:=1; i<=10; i++ {
	//fmt.Println(i)
	//}

	var myString = "GO!"
	for index, char := range myString {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	var slices = []string{"My ", "name ", "is ", "Dmitriy"}
	var strBuilder strings.Builder
	for i := range slices {
		strBuilder.WriteString(slices[i])
	}
	fmt.Println(strBuilder.String())
}
