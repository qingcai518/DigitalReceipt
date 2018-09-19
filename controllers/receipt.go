package controllers

import (
	"digitalReceipt/models"
	"github.com/astaxie/beego"
)

// Operations about Gifts
type ReceiptController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Receipts
// @Success 200 {object} models.Receipt
// @router / [get]
func (g * ReceiptController) GetAll() {
	println("1111111")
	receipts := models.GetAllReceipts()
	g.Data["json"] = receipts
	g.ServeJSON()
}
