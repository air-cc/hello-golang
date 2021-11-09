package dao

import (
	"database/sql"
	"fmt"
	"strings"

	"iaircc.com/go/demo/hello-gin/database"
	"iaircc.com/go/demo/hello-gin/model"
)

var db *sql.DB
var userStmtIns *sql.Stmt
var userStmtOut *sql.Stmt
var userStmtDel *sql.Stmt

func init() {
	var err error

	db = database.SqlDB

	// 预定义插入用户数据
	userStmtIns, err = db.Prepare("INSERT INTO t_users (name, age, sex) VALUES ( ?, ?, ? )")
	if err != nil {
		panic(err.Error())
	}

	// 预定义查询用户数据
	userStmtOut, err = db.Prepare("SELECT id, name, age, sex FROM t_users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	// 预定义删除用户数据
	userStmtDel, err = db.Prepare("DELETE FROM t_users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
}

// InsertUser 新增用户记录
func InsertUser(u model.User) sql.Result {
	result, err := userStmtIns.Exec(u.Name, u.Age, u.Sex)
	if err != nil {
		panic(err.Error())
	}

	return result
}

// FindUserByID 根据 id 查询用户记录
func FindUserByID(id int64) *model.User {
	var u = new(model.User)

	err := userStmtOut.QueryRow(id).Scan(&u.ID, &u.Name, &u.Age, &u.Sex)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return u
}

// UpdateUserByID 更新用户字段
func UpdateUserByID(u model.User) sql.Result {
	columValues := []string{}
	if u.Name != "" {
		columValues = append(columValues, fmt.Sprintf("name=%s", u.Name))
	}

	if u.Age != 0 {
		columValues = append(columValues, fmt.Sprintf("age=%d", u.Age))
	}

	if u.Sex != 0 {
		columValues = append(columValues, fmt.Sprintf("sex=%d", u.Sex))
	}

	columValuesStr := strings.Join(columValues, ",")

	sqlStr := fmt.Sprintf("UPDATE t_users SET %s WHERE id = %d", columValuesStr, u.ID)

	result, err := db.Exec(sqlStr)
	if err != nil {
		panic(err.Error())
	}

	return result
}

// DelUserByID 根据 id 删除用户记录
func DelUserByID(id int64) sql.Result {
	result, err := userStmtDel.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	return result
}
