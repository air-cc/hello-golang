package model

// User 用户
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
	Sex  uint8  `json:"sex" binding:"oneof=0 1"`
}
