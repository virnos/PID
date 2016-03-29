package main

import (
	"PID/models"
	_ "PID/routers"
	"fmt"

	"github.com/astaxie/beego"
)

func main() {

	fmt.Println("Starting....")

	fmt.Println("Start ok")
	models.Syncdb()
	beego.Run()
}
