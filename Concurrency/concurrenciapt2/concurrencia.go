package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func hi(num int) {
	fmt.Println("Hi", num)
	time.Sleep(1 * time.Second)
}

func get(num int) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(num))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil{
		panic(err)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go get(i)
	}

	var s string
	fmt.Scanln(&s)
}
