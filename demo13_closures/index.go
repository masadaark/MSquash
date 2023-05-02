package main

import "fmt"

//return func

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInt := intSeq()
	fmt.Println(newInt())
	newStr := myInitSeque()
	fmt.Println(newStr())
	fmt.Println(newStr())
	fmt.Println(myInitSeque()())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func myInitSeque() func() string {
	s := 0
	return func() string {
		s++
		return fmt.Sprintf("S = %d", s)
	}
}
