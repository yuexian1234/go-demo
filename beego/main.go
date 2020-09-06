package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:123456@(192.168.177.132:3306)/orm_test?charset=utf8")
}

func main() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
	profile := new(Profile)
	profile.Age = 30

	user := new(User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
}