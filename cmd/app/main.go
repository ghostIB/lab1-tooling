package main

import (
	"fmt"

	"github.com/ghostIB/lab1-tooling/internal"
)

func main() {
	res := internal.Add(5, 10)
	fmt.Printf("Sum: %d\n", res)

	// Навмисно не обробляємо помилку для лінтера (етап 3)
	val, _ := internal.Divide(10, 2)
	fmt.Printf("Result: %f\n", val)
}
