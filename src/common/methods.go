package common

import (
	"fmt"
)

// 函数和方法的区别
// 函数将变量作为参数：Function1(recv)
// 方法在变量上被调用：recv.Method1()
// 在接收者是指针时，方法可以改变接收者的值（或状态），这点函数也可以做到（当参数作为指针传递，即通过引用调用时，函数也可以改变参数的状态）。
// 方法没有和数据定义（结构体）混在一起：它们是正交的类型；表示（数据）和行为（方法）是独立的。

// 方法定义方法的一般格式如下
// func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
// 在方法名之前，func 关键字之后的括号中指定 receiver。
// 利用结构体和方法可以实现类对象访问方法

type TwoInts struct {
	a int
	b int
}

type ThreeInts struct {
	a int
	b int
	c int
}

// 类型和作用在它上面定义的方法必须在同一个包里定义，这就是为什么不能在 int、float 或类似这些的类型上定义方法
// 但是有一个间接的方式：可以先定义该类型（比如：int 或 float）的别名类型，然后再为别名类型定义方法。

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

// 不同类型的接收者可以重载
func (tn *ThreeInts) AddThem(param int) int {
	return tn.a + tn.b + tn.c + param
}

func TestReceiveMethods() {
	t2 := &TwoInts{1, 2}
	fmt.Printf("t2 is : %d \n", t2.AddThem())

	t3 := &ThreeInts{1, 2, 3}
	fmt.Printf("t3 is : %d \n", t3.AddThem(10))
}
