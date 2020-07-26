package main

import "fmt"

func main() {

	go func() {
		fmt.Println("hello world 1!")
		fmt.Println("hello world 1!")
	}()

	go func() {
		fmt.Println("hjjjj")
		fmt.Println("hjjjj")
		fmt.Println("hjjjj")
	}()

	fmt.Println("hello world 1!")
}
