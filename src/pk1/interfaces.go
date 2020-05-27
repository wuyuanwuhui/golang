package pk1

// 接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。
// 类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。
// 实现某个接口的类型（除了实现接口方法外）可以有其他的方法。
// 一个类型可以实现多个接口。接口类型可以包含一个实例的引用， 该实例的类型实现了此接口（接口是动态类型）。

import (
	"fmt"
)

type Shaper interface {
	Area() float32
	// AB() int // 如果类型没有实现此方法，则类型编译出错
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	width, length float32
}

// 接收者为实例拷贝
func (r Rectangle) Area() float32 {
	return r.width * r.length
}

func TestInterface() {
	sq := new(Square)
	sq.side = 5
	var areaIntf Shaper
	// 把square赋值给shaper接口类型的变量
	areaIntf = sq
	fmt.Printf("areaIntf area is : %v \n", areaIntf.Area())

	// 不同的接收者
	q := &Square{5}
	r := Rectangle{5, 6}
	shapers := []Shaper{q, r}
	fmt.Println("Looping through shapes for area ...")

	for i, _ := range shapers {
		fmt.Println("shaper is :", shapers[i])
		fmt.Println("Area of this shaper is :", shapers[i].Area())
	}

	// 一般方法调用接口类型方法
	var o valueable = stockPosition{"GODO", 3000, 3}
	// 会找到对应的类型方法调用
	showValue(o)
	o = car{"Bench", "Moe100", 1226}
	showValue(o)
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valueable interface {
	getValue() float32
}

func showValue(asset valueable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}
