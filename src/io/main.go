package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	firtName, lastName, s string
	i                     int
	f                     float32
	input                 = "56.12 / 5212 / Go"
	format                = "%f / %d /%s"
)

var inputReader *bufio.Reader
var inputs string
var err error

func main() {
	fmt.Println("Please enter your full name: ")
	//fmt.Scanln(&firtName, &lastName)
	// fmt.Scanf("%s, %s", &firtName, &lastName)
	// fmt.Printf("Hi %s %s \n", firtName, lastName)

	// fmt.Sscanf(input, format, &f, &i, &s)
	// fmt.Println("From the string we read: ", f, i, s)

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	inputs, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The inputs was: %s\n", inputs)
	}
}
