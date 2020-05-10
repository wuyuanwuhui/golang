// 独立的包,必须在src文件夹目录下
package common

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
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

func Abc() {
	fmt.Println("hello Abc")
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 { // x lies in the integer range
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}

func TestString() {
	fmt.Println("--------- testString begin --------- ")
	str := "hello world"

	// Contains
	isContain := strings.Contains(str, "hello")
	fmt.Printf("the str %s is contain 'hello' %t \n", str, isContain)

	// Index
	isIndex := strings.Index(str, "hello")
	fmt.Printf("the str %s is index 'hello' %d \n", str, isIndex)
	// strings.LastIndex(s, str string) int
	// 非 ASCII 编码的字符在父字符串中的位置 strings.IndexRune(s string, r rune) int

	// Replace
	// str = strings.Replace(str, "hello", "hi", -1)
	// fmt.Printf("The str has been replaced : %s \n", str)

	// Count
	fmt.Printf("The o letter has showed %d times \n", strings.Count(str, "o"))

	// Repeat
	fmt.Printf("To repeat x 5 times: %s \n", strings.Repeat("x", 5))

	//Tolower To Upper
	fmt.Printf("To lower HY: %s \n", strings.ToLower("HY"))
	fmt.Printf("To lower hy: %s \n", strings.ToUpper("hy"))

	// TrimSpace
	// strings.Trim(s, "cut") 来将开头和结尾的 cut 去除掉
	fmt.Printf("To trim abc : %s \n", strings.TrimSpace("    abc    "))

	// 分割字符串 strings.Fields(s) / Split
	strs := "a b c"
	fields := strings.Fields(strs)
	fmt.Println(fields)
	strs = "a1_b2_c3"
	fields = strings.Split(strs, "_")
	for _, val := range fields {
		fmt.Printf("- %s \n", val)
	}

	fmt.Println(fields)
	strs = strings.Join(fields, "") // _
	fmt.Println(strs)

	// strings.NewReader(str) 从字符串中读取内容
	// 与字符串相关的类型转换都是通过 strconv 包实现的。
	var orgstr string = "666"
	var an int
	var news string

	an, _ = strconv.Atoi(orgstr)
	fmt.Printf("The an int is %d \n", an)
	an += 10
	news = strconv.Itoa(an)
	fmt.Printf("The news is %s \n", news)

	fmt.Println("--------- testString end --------- ")
}

func TestTime() {
	fmt.Println("--------- testTime start --------- \n")
	t := time.Now()
	year := t.Year()
	month := t.Month()
	day := t.Day()
	hour := t.Hour()
	min := t.Minute()
	sec := t.Second()
	// var current string = year + "-" + month + "-" + day + " " + hour + ":" + min + ":" + sec
	fmt.Println(t)
	// fmt.Printf("current time is : %s \n", current)
	// if not be str convert
	fmt.Printf("current time is : %4d-%02d-%02d %d:%d:%d \n",
		year, month, day, hour, min, sec)

	// end
	fmt.Println("\n--------- testTime end --------- ")
}

// 加法
func AddMyNumber(a int, b int) int {
	var c int = 10
	c = c + (a + b)
	return c
}

// 类型转换
func ConvertTest() int {
	a := 5.0 // 只能在函数体内如此声明
	b := int(a)
	return b
}

// 获取环境变量
func FetchEnv() {
	var goos string = runtime.GOOS
	fmt.Println("The goos is :", goos)
	path := os.Getenv("PATH")
	fmt.Println("The path is :", path)
}

// 变量定义
func DefVar() {
	// test var define
	a, b, c := 1, 2, "abc"
	fmt.Println(fmt.Sprintf("a:%d, b:%d, c:%s", a, b, c))
	fmt.Println(str)
	fmt.Printf("the string length is %d \n", len(str))

	// test string
	//fmt.Printf("the first of the str is %s \n", str[0])
	str2 := " and str2"
	str3 := str + str2 // concat string 更好的办法是使用函数 strings.Join()
	fmt.Printf("The str3 is: %s \n", str3)
	hasPrefix := strings.HasPrefix(str3, "T")
	fmt.Printf("Does the str3 has prefix %t\n", hasPrefix)
	TestString()

	// test bool
	// fmt.Println(5 != 10) // true
	// fmt.Println(false && false)
	// fmt.Println(0x10)
	// d := 1e10
	// d = common.IntFromFloat64(d)
	// fmt.Println(d)
	// PrintNumbers()
	var ab byte = 'c'
	d := int(ab)
	fmt.Printf("%d\n", ab) // print 99
	fmt.Printf("%d\n", d)
}

func PrintNumbers() {
	fmt.Println("print Numbers begin: ")
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d \n", a)
	}
	fmt.Println("print Numbers end ")
}
