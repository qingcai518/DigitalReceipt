package models

import (
	"github.com/astaxie/beego/orm"
)

var (
	GiftList map[string]*Gift
)

func init() {
	orm.RegisterModel(new(Gift))
}

type Gift struct {
	 Id int `orm: column(id); pk`
	 Name string `orm: column(name)`
	 Detail string `orm: column(detail)`
	 Thumbnail string `orm: column(thumbnail)`
	 Price float64 `orm: column(price)`
}

func GetGift(id int) (g *Gift) {
	o := orm.NewOrm()
	var gift *Gift
	o.Raw("select * from gift where id =?", id).QueryRow(&gift)
	return gift
}

func GetAllGifts() []Gift {
	o := orm.NewOrm()
	var gifts []Gift
	o.Raw("select * from gift").QueryRows(&gifts)
	return gifts
}