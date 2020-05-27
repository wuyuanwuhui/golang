package common

import (
	"fmt"
	"strconv"
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

type Point struct {
	x, y float64
}

// 因为一个结构体可以嵌入多个匿名类型，所以实际上我们可以有一个简单版本的多重继承
// 就像：type Child struct { Father; Mother}

type ThreeInts struct {
	a     int
	b     int
	c     int
	Point // 匿名字段为内嵌结构体
}

func (p *Point) Abs() float64 {
	return p.x * p.y
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

// https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/10.6.md
// string 方法可以格式化对象类似于 toString
// 不要在 String() 方法里面调用涉及 String() 方法的方法
// 无限递归调用（TT.String() 调用 fmt.Sprintf 很快就会导致内存溢出
// 必须返回string
func (t2 *TwoInts) String() string {
	return "The t2.a is: " + strconv.Itoa(t2.a) + ", The t2.b is: " + strconv.Itoa(t2.b)
}

func TestReceiveMethods() {
	t2 := &TwoInts{1, 2}
	fmt.Printf("t2 is : %d \n", t2.AddThem())

	t3 := &ThreeInts{1, 2, 3, Point{1.1, 2.3}}
	abs := t3.Abs() // 达到继承的目的

	fmt.Printf("t3 is : %d \n", t3.AddThem(10))
	fmt.Printf("t3 abs is: %v \n", abs)

	// test to string
	fmt.Printf("The t2 is : %v \n", t2)
	fmt.Println(t2)
	fmt.Printf("two2 is: %T\n", t2)
	fmt.Printf("two2 is: %#v\n", t2)

	//cp := new(CameraPhone)
	cp := new(CameraPhone)
	fmt.Println(cp.Call())
	fmt.Println(cp.TakePicture())

	// v := new(Voodoo)
	// v.Magic()
	// v.MoreMagic()
}

// 多重继承: 在 Go 语言中，通过在类型中嵌入所有必要的父类型，可以很简单的实现多重继承

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring"
}

type Camera struct{}

func (c *Camera) TakePicture() string {
	return "Click"
}

type CameraPhone struct {
	Phone
	Camera
}

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}
