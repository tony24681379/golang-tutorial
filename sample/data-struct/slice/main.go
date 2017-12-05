package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)           //s == [2 3 5 7 11 13]
	fmt.Println("s[1:4] ==", s[1:4]) //s[1:4] == [3 5 7]

	// missing low index implies 0
	fmt.Println("s[:3] ==", s[:3]) //s[:3] == [2 3 5]

	// missing high index implies len(s)
	fmt.Println("s[4:] ==", s[4:]) //s[4:] == [11 13]
}
