package main

import (
	"fmt"
	"strconv"
)

func strConversion() (int, error) {
	num, err := strconv.Atoi("10")
	if err != nil {
		return 0, err
	}
	return num, nil
}

func main() {
	if num, err := strConversion(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}
