// cd project main.go
package main

import (
	"common"
	"fmt" // alias : fm "fmt"
	"runtime"
)

func main() {
	// fmt.Println("Hello World!")
	fmt.Printf("The golang version is %s \n", runtime.Version())
	// fmt.Println(AddMyNumber(1, 2))
	// fmt.Println(ConvertTest())
	// fmt.Println(VIOLET)

	// test env
	// common.FetchEnv()

	// test define variables
	// common.DefVar()

	// test date time print
	// common.TestTime()

	// test state
	// common.TestIf()

	// test func
	// common.TestFunc()

	// test switch
	// common.TestSwitch()
	common.TestFor()
}
