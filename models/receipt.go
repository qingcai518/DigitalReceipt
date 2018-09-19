package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Receipt))
}

type Receipt struct {
	Id int `orm:"column(id);pk"`
	ImagePath string `orm:"column(image_path)"`
	ReceiptAt string `orm:"column(receipt_at)"`
	Tel string `orm:"column(tel)"`
	TotalPrice float64 `orm:"column(total_price)"`
	AdjustPrice float64 `orm:"column(adjust_price)"`
	CreatedAt string `orm:"column(created_at)"`
	UpdateAt string `orm:"column(update_at)"`
}

func GetAllReceipts() []Receipt {
	o := orm.NewOrm()
	var receipts []Receipt
	o.Raw("select * from receipt").QueryRows(&receipts)
	return receipts
}