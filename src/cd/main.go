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
	// common.TestFor()

	// test add mul sub
	// add, mul, sub := common.TestAddMulSub(1, 6)
	// fmt.Printf("add: %d, mul: %d, sub: %d \n", add, mul, sub)
	// add, mul, sub = common.TestAddMulSub2(1, 6)
	// fmt.Printf("add: %d, mul: %d, sub: %d \n", add, mul, sub)

	// test minmax
	// a, b := common.TestMinMax(20, 3)
	// fmt.Printf("min: %d, max: %d \n", a, b)

	// test point
	// n := 0
	// reply := &n                              // why not * ?
	// common.TestPoint(10, 5, reply)           // 传递的是指向整形的变量指针
	// fmt.Printf("The reply is %d \n", *reply) // *reply是变量值

	// test 可变形参函数
	// slice := []int{1, 3, 5, 6, 7, 10}
	// // 如果参数被存储在一个 slice 类型的变量 slice 中，则可以通过 slice... 的形式来传递参数，调用变参函数。
	// common.TestLongParams(slice...)

	// test defer
	// common.TestDefer()

	// test pass by callback
	// common.TestPassFunc()

	// test closure
	// common.TestClosure()

	// test return by func
	// common.TestRetFunc()

	// test fibonacci
	// common.TestFibonacci()

	common.TestArrDefine()
}
