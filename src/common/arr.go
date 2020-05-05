package common

import (
	"fmt"
)

// 元素的数目，也称为长度或者数组大小必须是固定的并且在声明该数组时就给出
// 不同于PHP 数组的定义必须是有固定的长度，元素类型必须一致

func TestArrDefine() {
	// len(arr) is length
	// for or for range

	// var arr [5]int
	// arr = [5]int{1, 2, 3, 4, 5}

	// arr1 = new ([5]int) 的类型是 *[5]int，而  var arr2 [5]int 的类型是 [5]int
	// var arr = new ([5]int)

	// arr 定义与上面两种同样
	arr := [...]int{1, 2, 3, 4, 5}

	for _, val := range arr {
		fmt.Println(val)
	}
}
