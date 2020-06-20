package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var urls = []string{
	//"http://www.google.com/",
	"http://www.baidu.com/",
	"http://www.qq.com/",
}

func main() {
	// fmt.Println(urls)

	for _, url := range urls {
		resp, err := http.Head(url)

		if err != nil {
			fmt.Println("Error:", url, err)
		}

		fmt.Println(url, ": ", resp.Status)
	}

	res, err := http.Get("http://www.baidu.com/")
	if err != nil {
		fmt.Println("http Get err: ", err)
	}

	data, er := ioutil.ReadAll(res.Body)

	if er != nil {
		fmt.Println("readAll err: ", er)
	}

	fmt.Printf("Got: %q", string(data))
}
