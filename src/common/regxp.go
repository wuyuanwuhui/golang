package common

import (
	"fmt"
	"regexp"
	"strconv"
)

func TestRegExp() {
	//目标字符串
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]{2}"
	// 替换改变
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Math Found !")
	}

	// 更多方法中，必须先将正则模式通过 Compile 方法返回一个 Regexp 对象
	// 然后我们将掌握一些匹配，查找，替换相关的功能。
	re, _ := regexp.Compile(pat)
	//将匹配到的部分替换为"##.#"
	str := re.ReplaceAllString(searchIn, "**.*")
	fmt.Println(str)

	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}
