package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", Add(2, 3))
}

func Add(n int, m int) int {
	return n + m
}

