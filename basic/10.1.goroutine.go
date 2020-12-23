// go 协程

package main

import (
	"fmt"
	"time"
)

func sayHi(s string) {
	fmt.Println(s)
}

func main() {
	defer func() {
		fmt.Println("done")
	}()

	sayHi("hello")
	go sayHi("world")

	// 不加 sleep 程序直接结束, 不等协程处理完成
	time.Sleep(100 * time.Millisecond)
}
