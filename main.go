package main

import "fmt"

func Add(c, d int) int {
	return c + d
}

func main() {
	fmt.Println("CI/CD Demo: 1 + 2 =", Add(1, 2))
}
