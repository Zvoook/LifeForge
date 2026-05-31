package main

import (
	"errors"
	"fmt"
)

func main() {
	var a int = 1
	fmt.Println(a)

	var b, c = 3, 4
	fmt.Println(b, c)

	arr := [...]int32{1, 2, 3}
	fmt.Println(arr)

	d, e, err := devideandmod(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(d, e)
	}
}

func devideandmod(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("division by zero")
		return 0, 0, err
	}
	return a / b, a % b, err
}
