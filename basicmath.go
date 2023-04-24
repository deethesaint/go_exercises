package main

import "fmt"

func main() {
	var a float32
	var b float32

	fmt.Println("Enter a: ")
	fmt.Scanln(&a)
	fmt.Println("Enter b: ")
	fmt.Scanln(&b)

	fmt.Printf("Sum = %.2f\n", a+b)
	fmt.Printf("Diff = %.2f\n", a-b)
	fmt.Printf("Mult = %.2f\n", a*b)
	fmt.Printf("Quot = %.2f\n", a/b)

}
