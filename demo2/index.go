package main

import "fmt"

//function ทางเข้า
func main() {
	display()
	display2("i am handsome")
	display3("Mk", 2)
	const a, b = 2, 3
	fmt.Printf("value of %d+%d is %d\n", a, b, sum(a, b))
	var x, y int = swap(a, b)
	fmt.Printf("%d,%d=>%d,%d\n", a, b, x, y)
	_x, _y, title := swapV2(a, b)
	fmt.Printf("%d,%d=>%d,%d %s", a, b, _x, _y, title)
}

//ขึ้นต้นด้วยตัวเล็ก private function ตัวใหญ่ Public function
func display() {
	fmt.Println("testSnippet")
}

func display2(msg string) {
	fmt.Println(msg)
}

func display3(title string, version int) {
	fmt.Print(title)
	fmt.Println(version)
}

func sum(a int, b int) int {
	return a + b
}

//multiple return
func swap(a int, b int) (int, int) {
	return b, a
}

func swapV2(a int, b int) (x, y int, title string) {
	y = a
	x = b
	title = "Swap Function V2"
	return
}
