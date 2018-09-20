package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:BCTech_8888@/digital_receipt?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "root:longhash@/recika?charset=utf8")
	orm.SetMaxIdleConns("default", 100)
	orm.SetMaxOpenConns("default", 100)
}