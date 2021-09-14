package main

import (
	"fmt"
	"time"
)

var input string

func main() {
	t := time.Now()
	fmt.Println("Hello, World!", t.Format("2006-01-02 15:04:05"))
	time.Sleep(time.Second * 10)
	fmt.Scanf("%v", &input)
}
