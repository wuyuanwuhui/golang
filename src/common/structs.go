package common

import (
	"fmt"
	"strings"
)

type innerS struct {
	in1 int
	in2 int
}

// 首字母小写是私有类型
type person struct {
	age  int    "This is age tag" // Tag
	sex  string //
	name string // 如果需要外部能访问的话首字母大写
	// 在一个结构体中对于每一种数据类型只能有一个匿名字段
	innerS // 匿名字段, 初始化成员默认值为0
	// int		  // 匿名字段
}

func upPerson(p *person) {
	p.name = strings.ToUpper(p.name)
}

// 工厂方法创建结构体
func CreatePerson() *person {
	per := new(person)
	return per
}

// 导出字段
func (p *person) Age() int {
	return p.age
}

func (p *person) SetAget(age int) {
	p.age = age
}

func TestStructDefine() {
	fmt.Println("This is struct")
	var per1 person
	per1.age = 10
	per1.sex = "male"
	per1.name = "Yang"
	per1.in1 = 100
	per1.in2 = 200

	fmt.Println(per1)

	// t := new(T)，变量 t 是一个指向 T的指针，此时结构体字段的值是它们所属类型的零值。

	// 此时per2的类型是 *person 底层仍然会调用 new ()
	// 如果 File 是一个结构体类型，那么表达式 new(File) 和 &File{} 是等价的。
	// innerS 为内嵌结构体匿名字段
	per2 := &person{21, "female", "Yangmi", innerS{in1: 10, in2: 20}}
	fmt.Println(per2)

	// per3 := person{31, "male", "Zhang"}
	// per3 := new(person)
	// per3.age = 31
	// per3.sex = "male"
	// (*per3).name = "july"
	// fmt.Println(per3)

	// (*per3).name = "woodward" // 这是合法的
	// upPerson(per3)
	// fmt.Println(per3)

	// //pers3 := &Person{"Chris","Woodward"}
	// //upPerson(pers3)

	// // 使用工厂方法
	// per := CreatePerson()
	// per.age = 100
	// per.name = "Paul"
	// per.sex = "female"
	// fmt.Println(per)
}
