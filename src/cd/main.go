// cd project main.go
package main

import (
	fm "fmt" // alias
	"runtime"
)

const name = "Test name"

var str = "Test String"

type T struct{}

func main() {
	fm.Println("Hello World!")
	fm.Printf("%s", runtime.Version())
	fm.Println("")
	fm.Println(addMyNumber(1, 2))
}

func addMyNumber(a int, b int) int {
	var c int = 10
	c = c + (a + b)
	return c
}
