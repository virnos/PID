package models

import (
	"errors"
	"log"
	"os"
	"path"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	_ "github.com/mattn/go-sqlite3"

	"github.com/Unknwon/com"
)

const (
	_DB_NAME        = "data/PID.db"
	_SQLITE3_DRIVER = "sqlite3"
)

//设备表
type Project struct {
	Id            int64
	Username      string `orm:"size(32)" form:"Username"  valid:"Required;MaxSize(32);MinSize(4)"`
	Password      string `orm:"size(32)" form:"Password" valid:"Required;MaxSize(32);MinSize(6)"`
	IpAddress     string `orm:"size(32)" form:"IpAddress" valid:"Required;MaxSize(15);MinSize(7)"`
	Port          int16
	Path          string    `orm:"null;size(200)" form:"Path" valid:"MaxSize(200)"`
	Remark        string    `orm:"null;size(256)" form:"Remark" valid:"MaxSize(256)"`
	Status        int       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
	Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add" `
}

func (p *Project) TableName() string {
	return "Project" //beego.AppConfig.String("rbac_user_table")
}

func (p *Project) Valid(v *validation.Validation) {
	if len(p.Username) == 0 {
		v.SetError("Username", "用户名不允许为空")
	}
	//	if u.Password != u.Repassword {
	//		v.SetError("Repassword", "两次输入的密码不一样")
	//	}
}

//验证用户信息
func checkProject(p *Project) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&p)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func init() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Project))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)

}

/************************************************************/

//get user list
func QueryProject(page int64, page_size int64, sort string) (projects []orm.Params, count int64) {
	o := orm.NewOrm()
	project := new(Project)
	qs := o.QueryTable(project)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&projects)
	count, _ = qs.Count()
	return projects, count
}

//添加用户
func AddProject(p *Project) (int64, error) {
	if err := checkProject(p); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	//	project := new(Project)
	//	project.Username = u.Username
	//	project.Password = u.Password
	//	project.IpAddress = u.IpAddress
	//	project.Port = u.Port
	//	project.Remark = u.Remark
	//	project.Status = u.Status

	id, err := o.Insert(&p)
	return id, err
}

//更新用户
func UpdateProject(p *Project) (int64, error) {
	if err := checkProject(p); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	project := make(orm.Params)

	project["Username"] = p.Username

	project["Password"] = p.Password

	project["IpAddress"] = p.IpAddress

	project["Port"] = p.Port

	project["Remark"] = p.Remark

	project["Status"] = p.Status

	if len(project) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Project
	num, err := o.QueryTable(table).Filter("Id", p.Id).Update(project)
	return num, err
}

func DelProjectById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Project{Id: Id})
	return status, err
}

func GetProjectById(Id int64) (project Project) {
	project = Project{Id: Id}
	o := orm.NewOrm()
	o.Read(&project, "Id")
	return project
}
