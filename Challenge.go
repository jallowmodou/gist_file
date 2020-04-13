package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(sumOfSquares([]string{"3", "-1", "1", "14"}))
}

func sumOfSquares(numArr []string) int {
	i, err := strconv.Atoi(numArr[0])

	rest := numArr[1:]

	//Error checking
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return 0
	}

	square := i * i

	// negative & last number
	if i < 0 && len(rest) == 0 {
		return square
	}

	// negative & not last number
	if i < 0 && len(rest) > 0 {
		return sumOfSquares(rest)
	}

	// last man standing
	if i >= 0 && len(rest) == 0 {
		return square
	}

	return square + sumOfSquares(rest)

}


