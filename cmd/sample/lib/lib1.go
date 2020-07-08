package lib

import "fmt"

// Echo function
func Echo(t string) {
	fmt.Print(t)
}

// Roop function
func Roop() {
	s := []int{1, 2, 3}
	for _, i := range s {
		fmt.Print(i)
	}
}
