package main

import (
	"fmt"
	"reflect"
)

// 变量的最基本信息就是类型和值：反射包的 Type 用来表示一个 Go 类型，反射包的 Value 为 Go 值提供了反射接口。
// 两个简单的函数，reflect.TypeOf 和 reflect.ValueOf，返回被检查对象的类型和值

func main() {
	fmt.Println("reflect")

	var x float64 = 5.1
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println("the type of x :", t)
	fmt.Println("the value of x :", v)
	fmt.Println("the kind of x :", v.Kind())
	fmt.Println("the interface of x :", v.Interface())
}

// Go中变量类型以及方法的底层其实就是空接口 interface
