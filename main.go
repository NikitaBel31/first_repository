package main

import "fmt"

func main() {
	fmt.Println("Hello")
	a := make(chan int)
	go someFunc(a)
	fmt.Println(<-a)
	fmt.Println("Вот что вывелось")

}
func someFunc(a chan int) {
	a <- 12
}
