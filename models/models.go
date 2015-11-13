package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/go-sql-driver/mysql"
	//"strconv"
)

func init() {
}

type User struct {
	Id       int64  `orm:"auto"`
	Name     string `orm:"index;unique"`
	Password string `orm:"size(20)"`
	Phone    string `orm:"index;unique"`
	Total    int    `orm:"null"`
	Win      int    `orm:"null"`
}

func RegistDB() {
	//regist model
	_ = mysql.ErrBusyBuffer
	orm.RegisterModel(new(User))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "h00346577:2!h52112@tcp(182.254.241.198:3306)/studygolang?charset=utf8", 30, 30)
}

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
