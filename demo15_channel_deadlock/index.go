package main

import (
	"fmt"
	"time"
)

func main() {
	// errorRoutine() //error because has only one routine
	routineBuffer()
}

// func errorRoutine() {
// 	ch := make(chan int)
// 	ch <- 1 //send
// 	msg := <-ch
// 	fmt.Println(<-ch)
// 	fmt.Println(msg) //deadlock -> ไม่มีการจอง buffer และถูกบล็อก
// 	time.Sleep(1 * time.Second)
// }

func routineBuffer() {
	ch := make(chan int, 1) //ถังเก็บได้1ช่อง
	ch <- 1
	fmt.Println("step1")
	fmt.Println(<-ch) //ใช้ chanel รอบแรกก่อน
	ch <- 2
	fmt.Println("step2")
	fmt.Println(<-ch)
	time.Sleep(1 * time.Second)
}
