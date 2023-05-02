package main

import (
	"fmt"
	"time"
)

//sequential sync(blocking io)
//concurrent 1 server : many request (goroutine)
//parallel 2 server (switch request to cpu 1 or 2 (ex request คู่ เข้า 2 คี่ เข้า 1) a lot request)
/*
What is go-routines
-lightweight abstract thread => low level engine แตก งาน
-tiny stack size can resize as needed
-Start a go-routine using the go command: go f(atgs)
-the fo runtimes schedules go-routine on OS threads
-go-routines can communicate each other Channels
*/

// async
func main() {
	go run1()
	go run2()
	time.Sleep(5 * time.Second)
}

func run1() {
	for i := 0; i < 100; i++ {
		fmt.Println("Running1")
	}
}

func run2() {
	for i := 0; i < 100; i++ {
		fmt.Println("Running2")
	}
}
