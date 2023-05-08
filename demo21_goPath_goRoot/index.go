package main

import "gorm.io/gorm"

//gopath
//set env to run go.mod similar to package.json
//go auto import external lib

func main() {

}

type Product struct {
	gorm.Model
	Code   string
	Prince uint
}
