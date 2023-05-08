package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	go routine1(ch, 10)
	// waitForClose(ch)
	waitForClose1(ch)
}

func routine1(c chan int, countTo int) {
	for i := 0; i < countTo; i++ {
		c <- i
	}
	close(c)
}

func waitForClose(ch chan int) {
	for true {
		value, ok := <-ch
		if !ok {
			fmt.Println("Chanel is close")
			break
		}
		fmt.Println(value)
	}
}

func waitForClose1(ch chan int) {
	for value := range ch {
		fmt.Println(value)
	}
}
