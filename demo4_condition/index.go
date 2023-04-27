package main

import "fmt"

func main() {
	if result := doSomeThing(); result == true {
		fmt.Println("Wow")
	}
	fnSwitchCase("sss")
}

func doSomeThing() bool {
	return true
}

func fnSwitchCase(mode string) {
	switch mode {
	case "Mark":
		fmt.Println("found")
		break
	default:
		fmt.Println("some thing else")
	}
}
