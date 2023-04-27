package main

import "fmt"

func main() {
	// fncFor()
	// fnWhile()
	// fnForEach()
	whileBreak()
}

func fncFor() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Index %d\n", i)
	}
}

//no while in go land
func fnWhile() {
	index := 0
	for index < 5 {
		fmt.Printf("Index %d\n", index)
		index++
	}
}

func fnForEach() {
	arrName := []string{"mark", "sud", "tae"}
	for _, item := range arrName { //_ is void คือประกาศ index แล้วไม่ใช้
		fmt.Printf("%s item\n", item)
		if item == "sud" {
			break
		}
	}
}

func whileBreak() {
	index := 0
	for true {
		if index > 5 {
			break
		}
		fmt.Printf("Index %d\n", index)
		index++
	}
}
