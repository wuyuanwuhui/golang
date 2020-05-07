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

	// var arrLazy = []int{5, 6, 7, 8, 22}	//注：初始化得到的实际上是切片slice
	// [...]未知长度 ... 可同样可以忽略，从技术上说它们其实变化成了切片。
	var keyValues = [...]string{0: "hello", 5: "china"}
	fmt.Println(keyValues)
}

// 数组算法 fibonacci
func TestArrFabico() {
	var arr [10]int
	arr[0] = 1
	arr[1] = 1

	for i := 0; i < 10; i++ {
		if i > 1 {
			arr[i] = arr[i-1] + arr[i-2]
		}
	}

	fmt.Println(arr)
}

// 把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：
// 传递数组的指针、 使用数组的切片

func TestArrPt() {
	arr := [3]float32{1.2, 3.6, 5.5}
	sum := sumFloat(&arr) // to pass a pointer to the array
	fmt.Println("The sum of the arr is: %f \n", sum)
}

// 指向指针的数组
func sumFloat(arr *[3]float32) (sum float32) {
	for _, v := range arr {
		sum += v
	}
	return
}
