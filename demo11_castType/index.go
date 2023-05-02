package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var index int8 = 15
	var bigIndex int32
	bigIndex = int32(index)
	fmt.Println(bigIndex)
	lines_y := "50"
	yesterday, err := strconv.Atoi(lines_y) //try catch version go
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(yesterday)
}
