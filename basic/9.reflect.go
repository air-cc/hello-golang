// 反射

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Human 人类
type Human struct {
	name string
	age  int
}

func (h Human) String() string {
	return h.name + "-" + strconv.Itoa(h.age)
}

// Speak ...
func (h Human) Speak(message string) {
	fmt.Printf("%s: %s\n", h.name, message)
}

func main() {
	h := Human{"ss", 18}

	k := reflect.TypeOf(h)
	v := reflect.ValueOf(h)

	fmt.Println(k, v)

	fCount := v.NumField()

	for i := 0; i < fCount; i++ {
		field := k.Field(i)
		value := v.Field(i)

		fmt.Println(field.Name, value, reflect.TypeOf(value))
	}

	var a int = 1
	fmt.Println(reflect.TypeOf(a).Kind())
}
