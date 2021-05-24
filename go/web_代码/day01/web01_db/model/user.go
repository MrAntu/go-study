package model

import (
	"fmt"
	"web01_db/utils"
)

//User 结构体
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

//AddUser 添加User的方法一
func (user *User) AddUser() error {
	//1.写sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//2.预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常：", err)
		return err
	}
	//3.执行
	_, err2 := inStmt.Exec("admin", "123456", "admin@atguigu.com")
	if err2 != nil {
		fmt.Println("执行出现异常：", err2)
		return err
	}
	return nil
}

//AddUser2 添加User的方法二
func (user *User) AddUser2() error {
	//1.写sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//2.执行
	_, err := utils.Db.Exec(sqlStr, "admin2", "666666", "admin2@sina.com")
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}
