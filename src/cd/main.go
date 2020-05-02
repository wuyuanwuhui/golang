// cd project main.go
package main

import (
	fm "fmt" // alias
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// const use
const name = "Test name"
const Pi = 3.14159

type Color int

const (
	RED    Color = iota // 0
	ORANGE              // 1
	YELLOW              // 2
	GREEN               // ..
	BLUE
	INDIGO
	VIOLET // 6
)

// 全局变量定义
var age int
var isok bool = true
var str string = "Test String" // `Test String \n` 会被原样输出
var day = 1

type T struct{}

func main() {
	fm.Println("Hello World!")
	fm.Printf("%s", runtime.Version())
	fm.Print("\n")
	// fm.Println("")
	// fm.Println(addMyNumber(1, 2))
	// fm.Println(convertTest())
	// fm.Println(VIOLET)
	fetchEnv()
	defVar()
}

// 加法
func addMyNumber(a int, b int) int {
	var c int = 10
	c = c + (a + b)
	return c
}

// 类型转换
func convertTest() int {
	a := 5.0 // 只能在函数体内如此声明
	b := int(a)
	return b
}

// 获取环境变量
func fetchEnv() {
	var goos string = runtime.GOOS
	fm.Println("The goos is :", goos)
	path := os.Getenv("PATH")
	fm.Println("The path is :", path)
}

// 变量定义
func defVar() {
	// test var define
	a, b, c := 1, 2, "abc"
	fm.Println(fm.Sprintf("a:%d, b:%d, c:%s", a, b, c))
	fm.Println(str)
	fm.Printf("the string length is %d \n", len(str))

	// test string
	//fm.Printf("the first of the str is %s \n", str[0])
	str2 := " and str2"
	str3 := str + str2 // concat string 更好的办法是使用函数 strings.Join()
	fm.Printf("The str3 is: %s \n", str3)
	hasPrefix := strings.HasPrefix(str3, "T")
	fm.Printf("Does the str3 has prefix %t\n", hasPrefix)
	testString()

	// // test bool
	// fm.Println(5 != 10) // true
	// fm.Println(false && false)
	// fm.Println(0x10)
	// d := 1e10
	// // d = IntFromFloat64(d)
	// fm.Println(d)
	// // printNumbers()
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 { // x lies in the integer range
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fm.Sprintf("%g is out of the int32 range", x))
}

func printNumbers() {
	fm.Println("print Numbers begin: ")
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fm.Printf("%d \n", a)
	}
	fm.Println("print Numbers end ")
}

func testString() {
	fm.Println("--------- testString begin --------- ")
	str := "hello world"

	// Contains
	isContain := strings.Contains(str, "hello")
	fm.Printf("the str %s is contain 'hello' %t \n", str, isContain)

	// Index
	isIndex := strings.Index(str, "hello")
	fm.Printf("the str %s is index 'hello' %d \n", str, isIndex)
	// strings.LastIndex(s, str string) int
	// 非 ASCII 编码的字符在父字符串中的位置 strings.IndexRune(s string, r rune) int

	// Replace
	// str = strings.Replace(str, "hello", "hi", -1)
	// fm.Printf("The str has been replaced : %s \n", str)

	// Count
	fm.Printf("The o letter has showed %d times \n", strings.Count(str, "o"))

	// Repeat
	fm.Printf("To repeat x 5 times: %s \n", strings.Repeat("x", 5))

	//Tolower To Upper
	fm.Printf("To lower HY: %s \n", strings.ToLower("HY"))
	fm.Printf("To lower hy: %s \n", strings.ToUpper("hy"))

	// TrimSpace
	// strings.Trim(s, "cut") 来将开头和结尾的 cut 去除掉
	fm.Printf("To trim abc : %s \n", strings.TrimSpace("    abc    "))

	// 分割字符串 strings.Fields(s) / Split
	strs := "a b c"
	fields := strings.Fields(strs)
	fm.Println(fields)
	strs = "a1_b2_c3"
	fields = strings.Split(strs, "_")
	for _, val := range fields {
		fm.Printf("- %s \n", val)
	}

	fm.Println(fields)
	strs = strings.Join(fields, "") // _
	fm.Println(strs)

	// strings.NewReader(str) 从字符串中读取内容
	// 与字符串相关的类型转换都是通过 strconv 包实现的。
	var orgstr string = "666"
	var an int
	var news string

	an, _ = strconv.Atoi(orgstr)
	fm.Printf("The an int is %d \n", an)
	an += 10
	news = strconv.Itoa(an)
	fm.Printf("The news is %s \n", news)

	fm.Println("--------- testString end --------- ")
}
