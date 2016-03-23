package cms

import (
	m "PID/models/cms"

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
		this.TplName = "amazeui/rbac/user.tpl"
	}

}

func (this *ProjectController) AddProject() {
	p := m.Project{}
	if err := this.ParseForm(&p); err != nil {
		//handle error
		//		this.Rsp(false, err.Error())
		return
	}
	id, err := m.AddProject(&p)
	if err == nil && id > 0 {
		//		this.Rsp(true, "Success")
		return
	} else {
		//		this.Rsp(false, err.Error())
		return
	}

}

func (this *ProjectController) UpdateProject() {
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

func (this *ProjectController) DelProject() {
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
