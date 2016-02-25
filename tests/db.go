package main /*************************************************************************
   > File Name: db.go
   > Author: b0lu
   > Mail: b0lu_xyz@163.com
   > Created Time: 2016年02月25日 星期四 14时28分29秒
************************************************************************/
//package main1

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/zhangqin/sso-server/models"
)

func main() {
	//orm.RunCommand()
	// 数据库别名
	name := "default"

	// drop table 后再建表
	force := true

	// 打印执行过程
	verbose := true

	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
