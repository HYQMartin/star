package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"strconv"
	"time"
)

type User struct {
	Id       int64  `orm:"auto"`
	Name     string `orm:"index;unique"`
	Password string `orm:"size(256)"`
}

type Catalog struct {
	Id           int64
	Ident        string `orm:"unique"`
	Name         string
	Resume       string
	DisplayOrder int
	ImgUrl       string
}

type Blog struct {
	Id                    int64
	Ident                 string `orm:"unique"`
	Title                 string
	Keywords              string       `orm:"null"`
	CatalogId             int64        `orm:"index"`
	Content               *BlogContent `orm:"-"`
	BlogContentId         int64        `orm:"unique"`
	BlogContentLastUpdate int64
	Type                  int8 /*0:original, 1:translate, 2:reprint*/
	Status                int8 /*0:draft, 1:release*/
	Views                 int64
	Creator               string
	Created               time.Time `orm:"auto_now_add;type(datetime)"`
}

type BlogContent struct {
	Id      int64
	Content string `orm:"type(text)"`
}

type NewemployeeCatalog struct {
	Id           int64
	Ident        string `orm:"unique"`
	Name         string
	Resume       string
	DisplayOrder int
	ImgUrl       string
}

type NewemployeeBlog struct {
	Id                    int64
	Ident                 string `orm:"unique"`
	Title                 string
	Keywords              string                  `orm:"null"`
	CatalogId             int64                   `orm:"index"`
	Content               *NewemployeeBlogContent `orm:"-"`
	BlogContentId         int64                   `orm:"unique"`
	BlogContentLastUpdate int64
	Type                  int8 /*0:original, 1:translate, 2:reprint*/
	Status                int8 /*0:draft, 1:release*/
	Views                 int64
	Creator               string
	Created               time.Time `orm:"auto_now_add;type(datetime)"`
}

type NewemployeeBlogContent struct {
	Id      int64
	Content string `orm:"type(text)"`
}

func (*Catalog) TableEngine() string {
	return engine()
}

func (*Blog) TableEngine() string {
	return engine()
}

func (*BlogContent) TableEngine() string {
	return engine()
}

func engine() string {
	return "INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci"
}

func init() {
	orm.RegisterModelWithPrefix("bb_", new(Catalog), new(Blog), new(BlogContent))
	// orm.RegisterModelWithPrefix("newemployee_", new(CatalogNE), new(BlogNE), new(BlogContentNE))
	orm.RegisterModel(
		new(NewemployeeCatalog),
		new(NewemployeeBlog),
		new(NewemployeeBlogContent),
		new(CapabilityMap),
		new(Capabilities),
		new(TrainingSchedulePublish),
		new(TrainingScheduleCollect),
		new(User))
}

// func RegistDB() {
//regist model
// _ = mysql.ErrBusyBuffer
// orm.RegisterModel(new(User))
// orm.RegisterDriver("mysql", orm.DR_MySQL)
// orm.RegisterDataBase("default", "mysql", "h00346577:2!h52112@tcp(182.254.241.198:3306)/studygolang?charset=utf8", 30, 30)
// }
//
func AddUser() {

	result := RegistUser("martin", "123456")
	fmt.Println(result)

	check := CheckPassword("martin", "123456")
	fmt.Println(check)

	return
}

func RegistUser(name string, password string) string {
	o := orm.NewOrm()

	//验证合法性
	qs := o.QueryTable("user")
	var u User
	err := qs.Filter("Name", name).One(&u)
	if err != orm.ErrNoRows {
		return "exist such name"
	}
	user := new(User)
	user.Name = name
	user.Password = password
	o.Insert(user)
	return "success"
}

func CheckPassword(name string, password string) bool {
	o := orm.NewOrm()

	qs := o.QueryTable("user")
	var u User
	err := qs.Filter("Name", name).One(&u)
	if err != nil {
		if err == orm.ErrNoRows {
			return false
		}
	}
	if password == u.Password {
		return true
	}
	return false
}

func GetAllUser() ([]*User, error) {
	o := orm.NewOrm()
	users := make([]*User, 0)
	qs := o.QueryTable("user")
	_, err := qs.All(&users)
	return users, err
}
