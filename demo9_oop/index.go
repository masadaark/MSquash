package main

import "testGo/demo9_oop/employee"

func main() {
	e := employee.Init("sam", "sun", 30, 20)
	e.LeavesRemaining()
	x := employee.Init("sam", "run", 30, 20)
	x.LeavesRemaining()
}
