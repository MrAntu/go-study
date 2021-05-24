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
	_, err2 := inStmt.Exec("admin3", "123456", "admin3@atguigu.com")
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

//GetUserByID 根据用户的id从数据库中查询一条记录
func (user *User) GetUserByID() (*User, error) {
	//写sql语句
	sqlStr := "select id,username,password,email from users where id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, user.ID)
	//声明
	var id int
	var username string
	var password string
	var email string
	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		return nil, err
	}
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u, nil
}

//GetUsers 获取数据库中所有的记录
func (user *User) GetUsers() ([]*User, error) {
	//写sql语句
	sqlStr := "select id,username,password,email from users"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	//创建User切片
	var users []*User
	for rows.Next() {
		//声明
		var id int
		var username string
		var password string
		var email string
		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			return nil, err
		}
		u := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		users = append(users, u)
	}
	return users, nil
}
