package main

import "fmt"

func main() {
	arr := [...]int{2, 5, 3, 6, 1, 6, 4, 3, 9}

	mtrix := [][]int{
		{5, 5, 4},
		{1, 7, 2},
		{9, 6, 3},
	}

	sum := 0

	// for loop -> for init; condition; post
	// there is no kind of ++i :D
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println("Sum: ", sum)

	//nested for
	//just for example, print a 2d array in Go doesnt need to use nested for
	for i := 0; i < len(mtrix); i++ {
		fmt.Println(mtrix[i])
	}

	for i := 0; i < len(mtrix); i++ {
		for j := 0; j < len(mtrix[i]); j++ {
			fmt.Printf("%d ", mtrix[i][j])
		}
		fmt.Println()
	}

	// while loop
	count := 0
	for count < len(arr) {
		fmt.Print(count, " ")
		count++
	}
}
