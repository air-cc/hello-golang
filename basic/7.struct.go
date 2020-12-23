// 结构体

package main

import "fmt"

// MALE 男性
// FEMALE 女性
const (
	MALE = iota
	FEMALE
)

// Sex 性别
type Sex byte

// Human 人类
type Human struct {
	name string
	age  int
	sex  Sex
}

// Citizen 公民
type Citizen struct {
	Human
	country string
	id      int
}

// Employee 雇员
type Employee struct {
	Citizen
	id    int
	email string
}

// 绑定 method

func (h Human) speak(message string) {
	fmt.Printf("%s: %s\n", h.name, message)
}

// CallMe Human notify
func (h Human) CallMe(message string) {
	fmt.Printf("Hi %s, you got a message: %s\n", h.name, message)
}

// CallMe Employee notify
func (e Employee) CallMe(message string) {
	fmt.Printf("Hi %s(employeeId: %d), you got a message: %s\n", e.name, e.id, message)
}

// SetName 设置 Citizen 名字
func (c *Citizen) SetName(name string) {
	c.name = name // 实际是 (*c).name = name 但 GO 自动转换了
}

// SetName 设置 Employee 名字
func (e *Employee) SetName(name string) {
	e.name = name + "-E"
}

func main() {
	c := Citizen{Human: Human{name: "cc", age: 20, sex: FEMALE}, country: "cn", id: 1}

	fmt.Printf("Citizen: name %s, age: %d, country: %s, CitizenIdL %d\n", c.name, c.age, c.country, c.id)

	e := Employee{Citizen: Citizen{Human: Human{name: "ee", age: 18, sex: MALE}, country: "cn", id: 2}, id: 1, email: "ee@ee.com"}

	// 同名字段 默认使用最外层的字段
	fmt.Printf("Employee: name: %s, EmployeeId: %d, CitizenId: %d\n", e.name, e.id, e.Citizen.id)

	// struct 绑定 method
	e.CallMe("good job!")
	e.speak("hello world")

	c.SetName("ccNew") // 实际是 (&c).SetName("ccNew") 但 GO 自动转换了
	e.SetName("eeNew")
	fmt.Printf("set name: Citizen: %s, Employee: %s\n", c.name, e.name)
}
