package main

import "fmt"

//Pointer ไว้ชี้ตำแหน่งไม่ได้ไว้เก็บข้อมูล
func main() {
	msg := "message"
	var msgPointer *string = &msg //&คือการชี้ address

	fmt.Println(msgPointer)  //ออกมาเป็นเลขฐาน16 (ตำแหน่ง)
	fmt.Println(*msgPointer) //value
	changeMessage(&msg, "newMessage1")
	fmt.Println(msg) // newMessage1
	changeMessage(msgPointer, "newMessage2")
	fmt.Println(msg) // newMessage2
}

func changeMessage(aPointer *string, newMessage string) {
	*aPointer = newMessage //เปลี่ยนค่าของตำแหน่งตัวแปรที่เราชี้อยู่
}
