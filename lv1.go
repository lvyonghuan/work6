package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	username string
	password string
	question string
	secrecy  string
}

var username string    //用户名
var password string    //密码
var question string    //密保问题
var secrecy string     //答案
var newPassword string //新密码

func search(Case int) (bool, int) {
	switch Case {
	case 1: //注册时用户重名查询
		row, _ := db.Query("select username from user")
		var temp user
		for row.Next() {
			err := row.Scan(&temp.username)
			if err != nil {
				fmt.Println(err)
			}
			if temp.username == username {
				return false, 0
			}
		}
	case 2: //登录相关查询
		row, _ := db.Query("select username,password from user")
		var temp user
		for row.Next() {
			err := row.Scan(&temp.username, &temp.password)
			if err != nil {
				fmt.Println(err)
			}
			if temp.username == username {
				if temp.password == password {
					return true, 0 //返回0：就是返回0罢了，占位
				} else {
					return false, 1 //返回状态1：密码错误
				}
			}
		}
		return false, 2 //状态2：用户不存在
	case 3: //密保查询
		row, _ := db.Query("select username,secrecy from user")
		var temp user
		for row.Next() {
			err := row.Scan(&temp.username, &temp.secrecy)
			if err != nil {
				fmt.Println(err)
			}
			if temp.username == username {
				if temp.secrecy == secrecy {
					_, err := db.Exec("update user set password=? where username=?", newPassword, username)
					if err != nil {
						fmt.Println(err)
					}
					return true, 0 //0还是占位
				} else {
					return false, 1 //状态1：密保填写错误
				}
			}
		}
		return false, 2 //状态2：无此用户
	}
	return true, 0
}

func register(c *gin.Context) {
	username = c.PostForm("username")
	result, _ := search(1) //1表示注册查询
	if !result {
		fmt.Println("用户已注册")
		c.JSON(200, "用户已注册")
		return
	}
	password = c.PostForm("password")
	question = c.PostForm("question")
	secrecy = c.PostForm("secrecy")
	_, err := db.Exec("insert into user (username,password,question,secrecy) value (?,?,?,?)", username, password, question, secrecy)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("用户名：" + username + ",密码:" + password)
	c.JSON(200, "注册成功!")
}

func login(c *gin.Context) {
	username = c.PostForm("username")
	password = c.PostForm("password")
	result, Case := search(2) //2表示查阅用户是否注册以及用户名密码信息比对
	if result == true {
		fmt.Println("登陆成功")
		c.JSON(200, "登陆成功")
		return
	} else {
		if Case == 1 {
			fmt.Println("登陆失败，密码错误")
			c.JSON(200, "登陆失败，请检查密码是否正确")
			return
		} else if Case == 2 {
			fmt.Println("登陆失败，用户不存在")
			c.JSON(200, "登陆失败，未注册或用户名不正确")
			return
		}
	}
}

func Secrecy(c *gin.Context) {
	username = c.PostForm("username")
	secrecy = c.PostForm("secrecy")
	newPassword = c.PostForm("newPassword")
	result, Case := search(3)
	if result == false {
		if Case == 1 {
			fmt.Println("密保错误")
			c.JSON(200, "请检查密保填写是否正确")
			return
		} else if Case == 2 {
			fmt.Println("用户不存在")
			c.JSON(200, "未注册或用户名不正确")
			return
		}
	} else {
		fmt.Println("用户密码修改。用户名：" + username + "，新密码：" + newPassword)
		c.JSON(200, "密码修改成功！")
		return
	}
}

func main() {
	var dns = "root:0@tcp(127.0.0.1:3306)/user"
	db, _ = sql.Open("mysql", dns)
	r := gin.Default()
	r.GET("/register", register)
	r.GET("/login", login)
	r.GET("/secrecy", Secrecy)
	err := r.Run()
	if err != nil {
		return
	}
}
