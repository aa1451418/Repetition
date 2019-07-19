package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Pwd  string
}

type Article struct {
	Id       int    //`orm:"pk;auto"`
	ArtiName string //`orm:"size(20)"`
	Content  string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/repetition")
	orm.RegisterModel(new(User), new(Article))
	orm.RunSyncdb("default", false, true)
}
