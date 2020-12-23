package main

import "fmt"

// 简单
func add(x int, y int) int {
	return x + y
}

// 统一类型参数
func sub(x, y int) int {
	return x - y
}

// 多值返回
func swap(x, y string) (string, string) {
	return y, x
}

// 命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

// 不定数量参数
func morArg(arg ...string) {
	// arg 为 slice 类型
	for index, value := range arg {
		fmt.Println(index, value)
	}
}

// 传值
func incrValue(x int) int {
	x++

	return x
}

// 传指针
func incrPointer(x *int) int {
	*x = *x + 1

	return *x
}

// 传函数
type intFunc func(int) int

func funcArg(f intFunc, x int) int {
	return f(x)
}

// 变量定义
var i, j int = 1, 2

// 如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型
var c, python, java = 'c', false, i

func main() {
	a, b := swap("a", "b")
	x, y := split(17)
	fmt.Println(add(1, 2), sub(1, 2), a, b, x, y)
	fmt.Println(i, j, c, python, java)

	// 短声明变量, `:=` 结构不能使用在函数外
	// k := 3

	// 可变参函数
	morArg("a", "b", "c")

	// 传值
	val := 1
	val1 := incrValue(val)
	fmt.Println(val, val1)

	// 传指针
	val2 := incrPointer(&val)
	fmt.Println(val, val1, val2)

	// 传函数
	val3 := funcArg(incrValue, 10)
	fmt.Println(val3)
}
