package main

import (
	//	m "PID/models"
	_ "PID/routers"
	"os"
	"path"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"

	"github.com/Unknwon/com"
)

const (
	_DB_NAME        = "data/PID.db"
	_SQLITE3_DRIVER = "sqlite3"
)

func main() {
	registerDB()
	beego.Run()
}

func registerDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
}

//func initDB() {
//	project := m.Project{}
//	m.AddProject()
//}
