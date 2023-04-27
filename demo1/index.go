package main

import (
	"fmt"
)

// global variable
var count int = 0

func main() {
	// fmt.Println(str)
	// time.Sleep(1 * time.Second)
	// fmt.Println("Success")
	test2()
}

func test2() {
	//Explicit Declaration (static typeเปลี่ยนแปลง type ตัวแปลไม่ได้)
	var testV1 int = 0
	var testV2 string = "hello"
	const testV4 bool = false
	//implicit Declration
	testV3 := true
	fmt.Println(testV1)
	fmt.Println(testV2)
	fmt.Println(testV3)
	fmt.Println(testV4)
	count++
	fmt.Println(count)
}

func run() {

}
