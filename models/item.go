package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Item))
}

type Item struct {
	Id int `orm:"column(id);pk"`
	ReceiptId int `orm:"column(receipt_id)"`
	Name string `orm:"column(name)"`
	Price float64 `orm:"column(price)"`
	CreatedAt string `orm:"column(created_at)"`
	UpdateAt string `orm:"column(update_at)"`
}

func GetAllItems(receiptId int) []Item {
	o := orm.NewOrm()
	var items []Item
	o.Raw("select * from item where receipt_id = ?", receiptId).QueryRows(&items)
	return items
}

func AddItem(item Item) (itemID int64, err error) {
	o := orm.NewOrm()

	result, err := o.Raw("insert into item (receipt_id, name, price) values (?, ?, ?)", item.ReceiptId, item.Name, item.Price).Exec()

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}