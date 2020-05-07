package common

import (
	"fmt"
)

// map 是引用类型，可以使用如下声明：
// var map1 map[keytype]valuetype
// var map1 map[string]int
// 在声明的时候不需要知道 map 的长度，map 是可以动态增长的。未初始化的 map 的值是 nil。
// 数组可以视为一种简单形式的 map，key 是从 0 开始的整数
// 常用的 len(map1) 方法可以获得 map 中的 pair 数目，这个数目是可以伸缩的，因为 map-pairs 在运行时可以动态添加和删除。

func TestDefineMap() {
	// var m map[string]string
	// 必须初始化否则后续赋值编译出错
	// m := map[string]string{}
	// 不要使用 new，永远用 make 来构造 map
	m := make(map[string]string)
	m["id"] = "100"
	m["username"] = "Joe"
	m["sex"] = "female"

	fmt.Println(m)

	// for k, v := range m {
	// 	fmt.Println(k, ": ", v)
	// }

	m2 := make(map[int][]int)
	m2[1] = []int{1, 10}
	m2[2] = []int{2, 20}
	// m2[5] = []int{}

	fmt.Println(m2)
	// 测试键值是否存在
	// fmt.Println(m2[5] == nil) // true
	val1, isExist := m2[5]     // 用于测试键值是否存在
	fmt.Println(val1, isExist) // [], false
	delete(m2, 2)
	fmt.Println(m2)
}

// map类型的切片
func TestMapSlice() {
	items := make([]map[int]int, 5) // map类型的切片
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][10] = 1
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
}

// map 的排序 https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/08.5.md
// 如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序
