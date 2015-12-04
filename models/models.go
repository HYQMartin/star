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
	Password string `orm:"size(20)"`
	Phone    string `orm:"index;unique"`
	Total    int    `orm:"null"`
	Win      int    `orm:"null"`
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
		new(TrainingScheduleCollect))
}

//
//
//
//
//

// func RegistDB() {
//regist model
// _ = mysql.ErrBusyBuffer
// orm.RegisterModel(new(User))
// orm.RegisterDriver("mysql", orm.DR_MySQL)
// orm.RegisterDataBase("default", "mysql", "h00346577:2!h52112@tcp(182.254.241.198:3306)/studygolang?charset=utf8", 30, 30)
// }
//
func AddUser() {
	/*o := orm.NewOrm()
	user := new(User)
	user.Name = "martin"
	user.Password = "121212"
	user.Phone = "15813807799"
	user.Total = 1
	user.Win = 1

	a, err := o.Insert(user)
	fmt.Println("a", a)
	if err != nil {
		fmt.Println(err)
		return
	}
	return*/

	result := RegistUser("marfftin", "355", "15581368077699")
	fmt.Println("------------")
	fmt.Println(result)
	fmt.Println("-------***********-----")

	check := CheckPassword("marfftin", "355")
	fmt.Println(check)

	return
}

func RegistUser(name string, password string, phone string) string {
	o := orm.NewOrm()

	//验证合法性
	qs := o.QueryTable("user")
	var u User
	err := qs.Filter("Name", name).One(&u)
	if err != orm.ErrNoRows {
		return "exist such name"
	}
	err = qs.Filter("Phone", phone).One(&u)
	if err != orm.ErrNoRows {
		return "exist such phone"
	}
	user := new(User)
	user.Name = name
	user.Password = password
	user.Phone = phone
	o.Insert(user)
	return "success"
	//一个电话只能注册一次
	//一个名字只能注册一次
	//	o:=orm.NewOrm()
	//qs:=o.QueryTable("user")
	//	err:=qs.Filter().on
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

func AddTotalAndWin(name string, iswin bool) {
	o := orm.NewOrm()
	var u User
	o.QueryTable("user").Filter("name", name).One(&u, "Total", "Win")
	t := u.Total + 1
	w := u.Win + 1
	if iswin {
		_, err := o.QueryTable("user").Filter("name", name).Update(orm.Params{
			"total": t,
			"win":   w,
		})
		if err != nil {
			return
		}

	} else {
		_, err := o.QueryTable("user").Filter("name", name).Update(orm.Params{
			"total": t,
		})
		if err != nil {
			return
		}
	}

}
func AddWin(name string) {

}
