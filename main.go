package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

// AddIntToSlice adds i to s and returns it
func AddIntToSlice(i int, s []int) []int {
	s = append(s, i)
	return s
}
