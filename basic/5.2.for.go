package main

import "fmt"

func main() {
	sum := 0

	// for 循环
	for i := 0; i < 10; i++ {
		sum += i
	}

	sum2 := 1

	// 相当于 while
	for sum2 < 10 {
		sum2 += sum2
	}

	fmt.Println(sum, sum2)

	// 死循环
	// for {
	// 	fmt.Println("loop")
	// }

	m1 := map[string]string{"name": "cc", "pos": "coder"}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
