package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*User
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Name string `orm:"column(name);pk"`
	Password string `orm:"column(password)"`
	OwnerPubKey string `orm:"column(owner_pub_key)"`
	ActivePubKey string `orm:"column(active_pub_key)"`
	MemoPubKey string `orm:"column(memo_pub_key)"`
}

func AddUser(u User) string {
	return ""
}

func GetUser(name string) (u *User, err error) {
	o := orm.NewOrm()
	user := User{Name: name}
	error := o.Read(&user)
	if error!= nil {
		return nil, errors.New("User not exists")
	}

	return &user, nil
}

func GetAllUsers() map[string]User {
	o := orm.NewOrm()
	users := make(map[string]User)
	qs, error := o.QueryTable("user").All(users)

	println(error)
	println(qs)
	println(users)

	return users
}

//func GetAllUsers() map[string]*User {
//	o := orm.NewOrm()
//	users := make(map[string]*User)
//
//	qs, error := o.QueryTable("user").All(&users)
//
//	if error != nil {
//		println(error)
//	}
//
//	println(qs)
//	println(users)
//
//	return users
//}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Name != "" {
			u.Name = uu.Name
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) (a *User) {
	o := orm.NewOrm()
	var user User
	err := o.Raw("select * from user where name=? and password=md5(?)", username, password).QueryRow(&user)
	if err != nil {
		return nil
	}

	return &user
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
