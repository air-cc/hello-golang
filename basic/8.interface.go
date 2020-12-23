// 接口

package main

import (
	"fmt"
	"sort"
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

func (h Human) speak(message string) {
	fmt.Printf("%s: %s\n", h.name, message)
}

// Student ...
type Student struct {
	Human
	grade int
}

// 重载
func (s Student) speak(message string) {
	fmt.Printf("%s(%d): %s\n", s.name, s.grade, message)
}

// Students ...
type Students []Student

// 实现 sort.Interface 方法

// Len ...
func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less ...
func (s Students) Less(i, j int) bool {
	return s[i].grade < s[j].grade
}

// Employee ...
type Employee struct {
	Human
	position string
}

// 重载
func (s Employee) speak(message string) {
	fmt.Printf("%s(%s): %s\n", s.name, s.position, message)
}

// Speaker ...
type Speaker interface {
	speak(message string)
}

func main() {
	s := Student{Human{"ss", 18}, 6}
	fmt.Println(s)
	e := Employee{Human{"ee", 25}, "manager"}

	l := []Speaker{s, e}

	for _, speaker := range l {

		switch value := speaker.(type) {
		case Student:
			fmt.Println("this is a Student instance", value.name)
		case Employee:
			fmt.Println("this is a Employee instance", value.name)
		}

		// 同理与
		// if value, ok := speaker.(Student); ok {
		// 	fmt.Println("this is a Student instance", value.name)
		// }

		speaker.speak("ok")
	}

	s2 := Student{Human{"ss2", 8}, 1}
	sList := Students{s, s2}

	sort.Sort(sList)
	fmt.Println(sList)
}
