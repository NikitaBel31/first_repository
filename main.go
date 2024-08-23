package main

import "fmt"

func main() {
	fmt.Println("Hello")
	a := make(chan int, 2)
	go someFunc(a)
	fmt.Println(<-a)
	fmt.Println(<-a)
	close(a)
	fmt.Println(<-a)
	fmt.Println(<-a)
}
func someFunc(a chan int) {
	a <- 12
	a <- 10
	a <- 1
}
