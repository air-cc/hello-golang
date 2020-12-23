package main

import (
	"errors"
	"fmt"
)

/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
     // 代表一个Unicode码

float32 float64

complex64 complex128
*/

func main() {
	// 变量未定义是会赋予零值
	// 数值类型为 `0`
	// 布尔类型为 `false`
	// 字符串为 `""`（空字符串）
	var j int

	// 短声明变量, `:=` 结构不能使用在函数外
	i := 1
	f := float32(i)

	// 常量
	const a = "hello"

	// 错误类型
	err := errors.New("error-var")

	fmt.Println(j, i, f, a, err)
}
