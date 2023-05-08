package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	quit := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		// quit <- "end"
	}()
	task(c, quit)
}

// ตัวกระจายงานเลือกทำตัวที่มีงาน
func task(c, quit chan string) {
	for {
		select {
		case c <- "Running...":
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("waiting...")
			time.Sleep(50 * time.Microsecond)
		}
	}
}
