package users

import (
	"fmt"
	"reflect"
	"time"
)

// User 用户
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
	Sex  uint8  `json:"sex" binding:"oneof=0 1"`
}

// UserDao local storage
var UserDao = make(map[string]*User)

// Save 存储user
func Save(u User) (id string) {
	id = fmt.Sprintf("%d", time.Now().Unix())

	UserDao[id] = &User{
		ID:   id,
		Name: u.Name,
		Age:  u.Age,
		Sex:  u.Sex,
	}

	return
}

// Get 获取用户信息
func Get(id string) (User, bool) {
	user, ok := UserDao[id]
	if ok == false {
		return User{}, false
	}

	return *user, ok
}

// Update 更新用户信息
func Update(id string, u User) bool {
	userRecord, ok := UserDao[id]

	if ok == false {
		return false
	}

	fields := reflect.TypeOf(u)
	values := reflect.ValueOf(u)

	count := fields.NumField()

	for i := 0; i < count; i++ {
		field := fields.Field(i)
		value := values.Field(i)

		if value.IsZero() && field.Name != "Sex" {
			continue
		}

		switch field.Name {
		case "Name":
			userRecord.Name = value.String()
		case "Age":
			userRecord.Age = uint8(value.Int())
		case "Sex":
			userRecord.Sex = uint8(value.Int())
		default:
			continue
		}
	}

	return true
}

// Delete 删除用户信息
func Delete(id string) bool {
	_, ok := UserDao[id]
	if ok == false {
		return false
	}

	delete(UserDao, id)

	return true
}
