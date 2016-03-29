package controllers

import (
	m "PID/models"
	"fmt"

	"github.com/astaxie/beego"
)

type ProjectController struct {
	beego.Controller
}

func (this *ProjectController) Index() {
	page, _ := this.GetInt64("page")
	page_size, _ := this.GetInt64("rows")
	sort := this.GetString("sort")
	order := this.GetString("order")
	if len(order) > 0 {
		if order == "desc" {
			sort = "-" + sort
		}
	} else {
		sort = "Id"
	}
	projects, count := m.QueryProject(page, page_size, sort)
	if this.IsAjax() {
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &projects}
		this.ServeJSON()
		return
	} else {
		this.Data["projects"] = &projects
		this.TplName = "index.html"
	}

}
func (this *ProjectController) Add() {
	this.TplName = "add.html"
}
func (this *ProjectController) Save() {
	fmt.Printf("Save")
	p := m.Project{}
	if err := this.ParseForm(&p); err != nil {
		fmt.Printf(err.Error())
		//handle error
		//		this.Rsp(false, err.Error())
		return
	}
	id, err := m.AddProject(&p)
	fmt.Printf("AddProject.html")

	if err == nil && id > 0 {
		//		this.Rsp(true, "Success")
		fmt.Printf("index.html")
		this.Ctx.Redirect(302, "/project/index")
	} else {
		//		this.Rsp(false, err.Error())
		return
	}

}

func (this *ProjectController) Edit() {
	return
}

func (this *ProjectController) Update() {
	p := m.Project{}
	if err := this.ParseForm(&p); err != nil {
		//handle error
		//		this.Rsp(false, err.Error())
		return
	}
	id, err := m.UpdateProject(&p)
	if err == nil && id > 0 {
		//		this.Rsp(true, "Success")
		return
	} else {
		//		this.Rsp(false, err.Error())
		return
	}

}

func (this *ProjectController) Del() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelProjectById(Id)
	if err == nil && status > 0 {
		//		this.Rsp(true, "Success")
		return
	} else {
		//		this.Rsp(false, err.Error())
		return
	}
}
