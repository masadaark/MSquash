package main

import "fmt"

//defer stack first in lastout
func main() {
	defer fmt.Println("finally") //run after main
	defer fmt.Println("before finally")
	for i := 0; i < 2; i++ {
		defer fmt.Println("finish", i)
	}
	for i := 0; i < 10; i++ {
		fmt.Println("idx", i)
	}
}
