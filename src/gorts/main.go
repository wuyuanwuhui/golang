// gorts project main.go
package main

import (
	"common"
	"fmt"
)

func main() {
	fmt.Println("This is test goroutines with chan")
	// test goroutines chans
	common.TestGortsChan()
}
