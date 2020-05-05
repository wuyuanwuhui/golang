package common

import (
	"fmt"
	"strconv"
)

func TestIf() {
	fmt.Println("--------- test if start --------- \n")

	a := 10
	b := 5
	s := "hello"
	if a >= b {
		fmt.Println("a >= 5 \n")
		if a == b {
			fmt.Println("  a is equal b \n")
		} else {
			fmt.Println("  a != b \n")
		}
		if len(s) > 3 {
			fmt.Println("  len more than 3")
		} else {
			fmt.Println("  len less than 3")
		}
	} else {
		fmt.Println("a < 5 \n")
	}

	// if err := file.Chmod(0664); err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	fmt.Println("\n--------- test if end --------- ")
}

func TestFunc() {
	fmt.Println("--------- test func start --------- \n")
	s := "hello"
	ant, err := strconv.Atoi(s)
	// if err != nil { return err }
	fmt.Printf("the ant is %d, the err is '%s' \n", ant, err)
	result, ok := rtMultiValues(10, 20)
	fmt.Printf("the result is %d, the ok is %t \n", result, ok)
	fmt.Println("\n--------- test func end --------- ")
}

// 多值返回函数
func rtMultiValues(a int, b int) (sum int, ok bool) { // (int, bool)
	//return a + b, true
	sum = a + b
	ok = true
	return
}

func TestSwitch() {
	var authType int = 3

	// switch result := calculate()
	// switch a, b := x[i], y[j]
	switch authType {
	case 1:
		// case num1 > 0 && num1 < 10:
		fmt.Println("authType is 1")
		// fallthrough
		break
	case 2:
		fmt.Println("authType is 2")
		break
	case 3, 4:
		fmt.Println("authType is 3,4")
	default:
		fmt.Println("authType default")
		break
	}
}

func TestFor() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("---- the %d time ---- \n", i)
	// 	// for j := 0; j < 10; j++ {
	// 	// 	println(j)
	// 	// }
	// }
	// str := "Go is a beautiful language!"
	// fmt.Printf("The length of str is: %d\n", len(str))
	// for ix := 0; ix < len(str); ix++ {
	// 	fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	// }

	// 循环打印出尖子塔型字母
	// G
	// GG
	// GGG
	// 两个for循环
	// for i := 1; i <= 10; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		// if j == 3 {
	// 		// 	break
	// 		// }
	// 		fmt.Print("G")
	// 	}
	// 	fmt.Println("")
	// }

	// 一个for
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(strings.Repeat("G", i))
	// }

	// 不用内置方法
	// str := "G"
	// for i := 1; i <= 10; i++ {
	// 	println(str)
	// 	str += "G"
	// }

	// for range
	// str = "Hello World"
	// for key, val := range str {
	// 	fmt.Printf("key: %d, val %c, char %c \n", key, val, str[key])
	// }

	// 会导致无限循环
	// for i := 0; i < 3; {
	// 	fmt.Println("Value of i:", i)
	// }

	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}

	// break是退出某一层的循环体，本层循环后的代码不在执行 for语句和switch 语句用到
	// continue是退出本次循环不影响下一次循环, 后续代码继续执行 只能用于for
}
