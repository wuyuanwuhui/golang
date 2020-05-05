// 函数使用
package common

// 申明一个在外部定义的函数，你只需要给出函数名与函数签名，不需要给出函数体
// func flushICache(begin, end uintptr) // implemented externally

import (
	"fmt"
	//"strings"
	"time"
)

// 对外暴露的方法名称首字母一定大写，遵循驼峰法
// 命名返回值
func TestAddMulSub(a int, b int) (add int, mul int, sub int) {
	add, mul, sub = a+b, a*b, a-b
	return
}

// 非命名返回值
func TestAddMulSub2(a int, b int) (int, int, int) {
	return a + b, a * b, a - b
}

// 求最小最大
func TestMinMax(a int, b int) (int, int) {
	var c int  // temp
	if a > b { // exchange a, b
		c = a
		a = b
		b = c
	}
	return a, b
}

// no return 指针传递
func TestPoint(a, b int, reply *int) {
	// 改变reply的值
	// *reply is int
	*reply = a * b
}

// 如果列表参数类型不一致，则使用结构体来定义一个类型
// 1. 定义一个结构类型，假设它叫 Options，用以存储所有可能的参数
// type Options struct {
// 		par1 type1,
// 		par2 type2,
// ...
// }

// 2. 使用空接口: 如果一个变长参数的类型没有被指定，则可以使用默认的空接口 interface{}

// 传递变长参数
func TestLongParams(s ...int) { //a int, b int, c ...int 或者 s … interface{}
	if len(s) == 0 {
		fmt.Println("len s is 0")
	}
	for _, val := range s { // for range
		fmt.Println(val)
	}
}

// 关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数
// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）：
// 一直延迟延迟延迟 先进后出
func TestDefer() {
	defer hi("first")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d \n", i)
	}
	defer hi("last")
}

func hi(s string) {
	fmt.Println(s + " Hi")
}

// func func1(s string) (n int, err error) {
// 	defer func() {
// 		log.Printf("func1(%q) = %d, %v", s, n, err)
// 	}()
// 	return 7, io.EOF
// }

// func main() {
// 	func1("Go")
// }

func TestPassFunc() {
	a := "this is for hi callback -- "
	showWithHi(a, hi) // callback 只需要带函数名即可
	//strings.IndexFunc()
}

// 传递参数为函数相当于callback
func showWithHi(a string, f func(string)) {
	f(a)
}

// closure
func TestClosure() {
	f := func(i int) int {
		return i * 10
	} // (v)表示自调用并且使用外部变量v
	r := f(100)
	fmt.Println(r)
}

func TestRetFunc() {
	start := time.Now()
	a1 := add1()
	b := 1
	fmt.Printf("The add1: %d \n", a1(b))
	a2 := add2(5)
	fmt.Printf("The add2: %d \n", a2(b))
	fmt.Printf("The add2: %d \n", a2(b))

	adr := Adder()
	fmt.Printf("The adr: %d \n", adr(b)) // 1
	fmt.Printf("The adr: %d \n", adr(b)) // 2
	fmt.Printf("The adr: %d \n", adr(b)) // 3
	end := time.Now()
	times := end.Sub(start)

	time.Sleep(10)
	fmt.Printf("it took this amount of time: %s\n", times)
}

// 应用闭包：将函数作为返回值
func add1() func(b int) int { // 返回值为函数
	return func(b int) int { // 闭包函数
		return b + 1
	}
}

// 带参数
func add2(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

// https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/06.9.md
// x 在多次调用中，变量 x 的值是被保留的
// 闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量
// 在闭包中使用到的变量可以是在闭包函数体内声明的，也可以是在外部函数声明的
func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

const LIM = 10

var fibs [LIM]uint64

func TestFibonacci() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

// fibonacci
func fibonacci(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}
