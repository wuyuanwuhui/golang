package common

import (
	"fmt"
	"sort"
)

// 优点 因为切片是引用，所以它们不需要使用额外的内存并且比使用数组更有效率
// 由起始和终止索引标识的一些项的子集。需要注意的是，终止索引标识的项不包括在切片内
// 切片是一个 长度可变的数组
// 声明切片的格式是： var identifier []type（不需要说明长度）。s := []int{1,2,3}。
// 一个切片在未初始化之前默认为 nil，长度为 0。
// 切片的初始化格式是：var slice1 []type = arr1[start:end]。
// 注意 绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针!!

func TestSlice() {
	var arr1 [6]int
	arr1 = [6]int{10, 11, 12, 13, 14, 15}

	// 这表示 slice1 是由数组 arr1 从 start 索引到 end-1 索引之间的元素构成的子集
	// start 代表下表起始，end代表第几个结束
	var slice1 []int = arr1[2:5] // [start:end] item at index 5 not included!

	fmt.Println(arr1)
	fmt.Println(slice1)
	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// // grow the slice
	// slice2 := slice1[0:4] // [12 13 14 15] 切片会向后移
	// fmt.Println(slice2)
	// fmt.Printf("The length of slice2 is %d\n", len(slice2))
	// fmt.Printf("The capacity of slice2 is %d\n", cap(slice2))

	// slice2 = slice2[0:5] // 越界 编译出错 slice1 2 永远不能超过数组长度
	// // 如此定义就是slice
	// b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	// bs := b[1:4]
	// fmt.Println("byte slice: ")
	// fmt.Println(bs) // [111 108 97] 输出了ascii
	// var s5 []int
	// fmt.Println(cap(s5))
	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("The sum is :", passBySlice(arr2[:])) // 注意传递的是切片arr2[:]
}

// 将切片传递给函数
func passBySlice(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

// 用 make() 创建一个切片
// 当相关数组还没有定义时，我们可以使用 make() 函数来创建一个切片
// 同时创建好相关数组：var slice1 []type = make([]type, len)。
// 也可以简写为 slice1 := make([]type, len)，这里 len 是数组的长度并且也是 slice 的初始长度。
// s := make([]int, 10, 20)

func TestMakeSlice() {
	// s := make([]byte, 5)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	// s = s[2:4]

	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Println(s1)
	fmt.Println(s2)
	s2[1] = 't'

	fmt.Println(s1)
	fmt.Println(s2)
}

// for range 可以遍历多维数组和切片
func TestChangeSlice() {
	s := [5]int{1, 2, 3, 4, 5}
	fmt.Println(s)
	for i := range s {
		s[i] = s[i] * 2
	}
	fmt.Println(s)
	fmt.Println(s[0:1])
}

func TestCopyAppendSlice() {
	// slFrom := []int{1, 2, 3}
	// slTo := make([]int, 10)

	// n := copy(slTo, slFrom)
	// fmt.Println(slTo)
	// fmt.Printf("Copied %d elements\n", n) // n == 3

	// sl3 := []int{1, 2, 3}
	// sl3 = append(sl3, 4, 5, 6)
	// fmt.Println(sl3)

	// b := []byte{'a', 'b', 'c'}
	// fmt.Println(b)
	// for _, v := range b {
	// 	// fmt.Printf("v %c \n", string(v))
	// 	fmt.Println(string(v))
	// }

	s := "hello"
	c := []byte(s) // 来获取一个字节的切片 c
	fmt.Println(string(c))
	// var b []byte
	// b = append(b, s)
	s2 := s[1:3] // 截取字符串
	fmt.Println(s2)
	c[0] = 'f'    // cannot use "f" // 不能通过s[0] = 'f'直接改变s
	s = string(c) // s 被改变
	fmt.Println(s)
}

func TestSort() {
	// 标准库提供了 sort 包来实现常见的搜索和排序操作
	arr := []int{1, 10, 2, 6, 7}
	s := arr[:]
	// 不能直接传递arr
	sort.Ints(s) // 编译类型的语言真不一样，如果你就这样定义没有调用到他会编译出错不同于PHP
	// func Float64s(a []float64) 来排序 float64 的元素
	// func Strings(a []string) 排序字符串元素。
	// func SearchInts(a []int, n int) int 进行搜索，并返回对应结果的索引值 但先要被排序

	//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/07.6.md
	//append 函数常见操作
	// 将切片 b 的元素追加到切片 a 之后：a = append(a, b...)
	var a []int = make([]int, 3)
	a = append(a, 5, 1, 3)
	fmt.Println(a)
	var b = make([]int, 6)
	// copy slice
	copy(b, a)
	// a = append(a, b) // can not append
	fmt.Println(b)
	// 删除位于索引 i 的元素：a = append(a[:i], a[i+1:]...)
	// 切除切片 a 中从索引 i 至 j 位置的元素：a = append(a[:i], a[j:]...)
	// 为切片 a 扩展 j 个元素长度：a = append(a, make([]T, j)...)
	// 在索引 i 的位置插入元素 x：a = append(a[:i], append([]T{x}, a[i:]...)...)
	// 在索引 i 的位置插入长度为 j 的新切片：a = append(a[:i], append(make([]T, j), a[i:]...)...)
	// 在索引 i 的位置插入切片 b 的所有元素：a = append(a[:i], append(b, a[i:]...)...)
	// 取出位于切片 a 最末尾的元素 x = a[len(a)-1]
	// 元素 x 追加到切片 a：a = append(a, x)
	a = append(a, b...)
	fmt.Println(a)
}
