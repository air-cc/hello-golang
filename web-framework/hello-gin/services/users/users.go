package users

import (
	"iaircc.com/go/demo/hello-gin/dao"
	"iaircc.com/go/demo/hello-gin/model"
)

// Save 存储 user
func Save(u model.User) (id int64) {
	result := dao.InsertUser(u)

	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	return
}

// Get 获取用户信息
func Get(id int64) *model.User {
	u := dao.FindUserByID(id)

	return u
}

// Update 更新用户信息
func Update(u model.User) bool {
	result := dao.UpdateUserByID(u)
	af, err := result.RowsAffected()
	if err != nil {
		return false
	}

	return af > 0
}

// Delete 删除用户信息
func Delete(id int64) bool {
	result := dao.DelUserByID(id)
	af, err := result.RowsAffected()
	if err != nil {
		return false
	}

	return af > 0
}
