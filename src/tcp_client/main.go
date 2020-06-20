package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	// connect to server
	// testConnectToServer()

	// socket
	testUseSocketConn()
}

func testConnectToServer() {
	// 创建连接
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")

	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}
}

func testUseSocketConn() {
	var (
		host          = "localhost"
		port          = "80"
		remote        = host + ":" + port
		msg    string = "GET / \n"
		data          = make([]uint8, 4096)
		read          = true
		count         = 0
	)
	// create socket
	con, err := net.Dial("tcp", remote)
	//
	io.WriteString(con, msg)
	//
	for read {
		count, err = con.Read(data)
		if err != nil {
			fmt.Println("Error read data ", err.Error())
		}
		read = (err == nil)
		fmt.Printf(string(data[0:count]))
	}
}
