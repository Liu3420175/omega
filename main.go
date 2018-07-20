package main

import (
	_ "omega/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"omega/conf"
	"fmt"
    _ "omega/auth"
)
func CreateTable(){
	name := "default"
	err := orm.RunSyncdb(name,false,true)
	if err != nil{
		beego.Error(err)
	}
}

func main() {

	databases := conf.DATABASES
	database := databases["default"]
	orm.RegisterDriver("mysql",orm.DRMySQL)
	connent := database["USER"] + ":" + database["PASSWORD"] + "@tcp(" + database["HOST"] + ":" + database["PORT"] + ")/" + database["NAME"] + "?charset=utf8"
	fmt.Println(connent)
	orm.RegisterDataBase("default","mysql",connent)
	//o.Using("default")
	CreateTable()
	beego.Run()
}
