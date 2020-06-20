package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	inputFile, inputError := os.Open("data.txt")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}

	defer inputFile.Close()

	// 循环读取文件
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readError := inputReader.ReadString('\n')
		fmt.Printf("The input was %s ", inputString)
		if readError == io.EOF {
			fmt.Println("\n----------\n")
			// return
			break
		}
	}

	// 一次性读取文件
	inputFile2 := "data.txt"
	outputFile := "data2.txt"
	buf, err := ioutil.ReadFile(inputFile2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s \n", err)
	}
	fmt.Printf("ioutil.ReadFile from inputFile2 is:\n%s  \n", string(buf))

	err = ioutil.WriteFile(outputFile, buf, 0644)

	if err != nil {
		panic(err.Error())
	}
}
