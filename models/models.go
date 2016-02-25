/*************************************************************************
   > File Name: models/models.go
   > Author: b0lu
   > Mail: b0lu_xyz@163.com
   > Created Time: 2016年02月25日 星期四 14时05分02秒
************************************************************************/
package models

import (
	"github.com/astaxie/beego/orm"     //引入beego的orm
	_ "github.com/go-sql-driver/mysql" //引入beego的mysql驱动
)

type Users struct {
	Id       int
	Username string `orm:"size(30)"`
	Password string `orm:"size(30)"`
	Phone    string `orm:"size(15)"`
}

func (u *Users) CheckLogin(username, password string) bool {
	flag := true
	o := orm.NewOrm()
	r := o.QueryTable("sso_users").Filter("username", username).Filter("password", password).One(u)
	if r != nil {
		flag = false
	}
	return flag
}

func (u *Users) GetUserInfoById(uid int) {
	o := orm.NewOrm()
	o.QueryTable("sso_users").Filter("id", uid).One(u)
}

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/sso?charset=utf8")
	orm.SetMaxIdleConns("default", 30) //设置数据库最大空闲连接
	orm.SetMaxOpenConns("default", 30) //设置数据库最大连接数
	orm.RegisterModelWithPrefix("sso_", new(Users))
}
