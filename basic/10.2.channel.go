// // 信道

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sum(a []int, c chan int) {
// 	total := 0

// 	for _, v := range a {
// 		total += v
// 	}

// 	fmt.Println("sum", total)

// 	c <- total
// }

// func fibonacci(n int, c chan int) {
// 	x, y := 1, 1

// 	for i := 0; i < n; i++ {
// 		c <- x
// 		fmt.Println("producer", x)
// 		x, y = y, x+y
// 	}

// 	close(c)
// }

// func simple() {
// 	// 无缓冲channel
// 	a := []int{1, 2, 3, 4, 5}

// 	ci := make(chan int)

// 	go sum(a[:2], ci)
// 	go sum(a, ci)
// 	x := <-ci
// 	y := <-ci
// 	fmt.Println(x, y)
// }

// func numLimit() {
// 	// 限定个数的channel
// 	ci2 := make(chan int, 2)
// 	ci2 <- 1
// 	ci2 <- 2
// 	x2 := <-ci2 // #1
// 	ci2 <- 3    // 无 #1 行读取操作，则报错：all goroutines are asleep - deadlock!

// 	fmt.Println(x2)
// }

// func useRange() {
// 	// 更多用法
// 	fmt.Println("useRange")
// 	ci3 := make(chan int, 4)
// 	go fibonacci(cap(ci3), ci3)

// 	// 生产端需要手动的 close channel 否则 调用 range 会报错
// 	for i := range ci3 {
// 		fmt.Println("customer", i)
// 	}

// 	fmt.Println("useRange done!")
// }

// func useSelect() {
// 	ci := make(chan int)
// 	cs := make(chan string)

// 	go func() {
// 		fmt.Println("set value")
// 		ci <- 1
// 		cs <- "a"
// 	}()

// 	// i = <-ci
// 	// fmt.Println(i)

// 	// 等待
// 	for {
// 		select {
// 		case i := <-ci:
// 			fmt.Println("i", i)
// 		case s := <-cs:
// 			fmt.Println("s", s)
// 		case <-time.After(5 * time.Second):
// 			fmt.Println("time up")
// 			break
// 		default:
// 			fmt.Println("no channel")
// 		}
// 	}
// }

// func main() {
// 	useSelect()
// }

package main

import (
	"fmt"
	"time"
)

func wait4OneSec(ch chan int, i int) {
	fmt.Println("wait4OneSec", i)

	time.Sleep(time.Second)

	ch <- i
}

func runUntil(ch chan int, max int) {
	i := 0
	len := cap(ch)

	for i < len {
		ch <- i

		if i >= max {
			close(ch)
			break
		}

		i++
	}

	close(ch)
}

func wait4OneSecAndClose(ch chan int, i int) {
	fmt.Println("wait4OneSec", i)

	time.Sleep(time.Second)

	ch <- i
	close(ch)
}

func main() {
	// 普通 channel
	fmt.Println("<<<< basic >>>>")

	ci := make(chan int)
	go wait4OneSec(ci, 0)
	fmt.Println(<-ci)

	// 带缓冲的 channel
	fmt.Println("<<<< buffer >>>>")

	cBuffer := make(chan int, 2)
	go wait4OneSec(cBuffer, 1)
	go wait4OneSec(cBuffer, 2)
	fmt.Println(<-cBuffer)
	fmt.Println(<-cBuffer)

	// 主动关闭 channel
	fmt.Println("<<<< range & close >>>>")

	cClose := make(chan int)
	go wait4OneSecAndClose(cClose, 3)
	fmt.Println(<-cClose)
	i, ok := <-cClose
	fmt.Println(i, ok) // 0 false

	cRange := make(chan int, 3)
	go runUntil(cRange, 4)

	// 循环 for i := range c 会不断从信道接收值，直到它被关闭
	for i := range cRange {
		fmt.Println(i)
	}

	// select channel
	fmt.Println("<<<< select >>>>")

	// select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("2s")
	case <-time.After(time.Second):
		fmt.Println("1s")
	default:
		// fmt.Println("not ready")
	}
}
