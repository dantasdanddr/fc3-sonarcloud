package main

import "fmt"

func main() {
	fmt.Println("Olá, mundo!")
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}

func sumX(a int, b int) int {
	return a + b + a
}
