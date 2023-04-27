package main

import "fmt"

//dictionary on golang
//key must be unique
func main() {
	var colors = map[string]int{"red": 1, "green": 2, "blue": 3} //static Map
	fmt.Println(colors)
	var colors2 = make(map[string]string) //blank map
	colors2["red"] = "#f00"
	colors2["green"] = "#0f0"
	fmt.Println(colors2)

	var courses = make(map[string]map[string]int) //2dMap
	courses["Android"] = map[string]int{"price": 200}
	courses["ios"] = make(map[string]int)
	courses["ios"]["price"] = 400
	fmt.Println(courses)
}
