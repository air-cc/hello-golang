package main

import (
	"fmt"
	"runtime"
	"time"
)

func getOS() string {
	yourOS := "Your OS is "

	// switch 不需要显示的 break 默认执行完 case 就退出
	switch os := runtime.GOOS; os {
	case "darwin":
		yourOS += "MacOS"
	case "linux":
		yourOS += "linux"
	default:
		yourOS += os
	}

	return yourOS
}

func main() {

	fmt.Println(getOS())

	t := time.Now()
	hour := t.Hour()

	fmt.Println(hour)

	// 没有条件的 switch
	// 这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() == 12:
		fmt.Println("nooning")
	default:
		fmt.Println("afternoon")
	}

	// 一个 case 中匹配多个条件
	i := 2

	switch i {
	case 0:
		fmt.Println("0")
	case 1, 2, 3:
		fmt.Println("1,2,3")
		fallthrough // switch 默认 break, 通过 fallthrough 强制执行下一个 case
	default:
		fmt.Println("no match")
	}
}
