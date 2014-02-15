package main

import "fmt"

func main() {
	fmt.Println("2 + 3 = %v", Add(2, 7))
}

func Add(n int, m int) int {
	return n + m
}

