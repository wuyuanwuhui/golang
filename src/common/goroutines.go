package common

import (
	"fmt"
	"time"
)

func TestGortsChan() {
	ch := make(chan string)
	defer close(ch)
	// 2个协程启动完成后，并行运行，输出无顺序 每次都不一样
	go sendData(ch)
	go getData(ch)
	printTen()

	time.Sleep(1e9) // 让出启动协程时间为1s完成启动
}

func sendData(ch chan string) {
	ch <- "Was"
	ch <- "Beijing"
	ch <- "Chengdu"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("get : %s \n", input)
	}
}

func printTen() {
	for i := 1; i <= 200; i++ {
		fmt.Println(i)
	}
}
