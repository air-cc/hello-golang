// 复合类型

package main

import "fmt"

func main() {
	// 数组
	arr1 := [3]int8{1, 2, 3}
	arr2 := [...]int{1, 2, 3, 4, 5}
	arr3 := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("array:", arr1[0], arr2[4], arr3[1][0])

	// slice - 引用类型
	silce1 := []int{}
	silce2 := []int{1, 2, 3}
	silce3 := arr2[1:3]
	silce4 := arr2[:]

	silce3[0] = 22

	fmt.Println("slice:", silce1, silce2[2], silce3[0], silce4[1], len(silce4))

	// map - 引用类型
	m1 := map[string]int{"c": 1, "go": 2}
	m2 := make(map[string]int)
	m2 = m1

	delete(m1, "c")
	info, ok := m2["c"]

	// 获取成功
	if ok {
		fmt.Println("get c", info)
	} else {
		fmt.Println("no c info found")
	}

	m3 := make(map[string]int)
	fmt.Println("m3", m3)

	/**
		- make 用于给 slice / map / channel 这几个内建复合类型分配内存, 初始化变量 并返回引用
		- new 用于给其他自定义类型分配内存，未初始化变量，值为 零 值，返回指针
	**/

}
