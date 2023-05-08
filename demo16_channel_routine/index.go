package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	go routine1(ch, 1) //no blocking main
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) //over read will be deadlock
	time.Sleep(1 * time.Second)
}

func routine1(c chan int, payload int) {
	c <- payload
	fmt.Println("s1")
	c <- payload
	fmt.Println("s2")
	c <- payload
	fmt.Println("s3")
}
