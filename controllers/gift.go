package controllers

import (
	"digitalReceipt/models"

	"github.com/astaxie/beego"
)

// Operations about Gifts
type GiftController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Gifts
// @Success 200 {object} models.Gift
// @router / [get]
func (g * GiftController) GetAll() {
	println("1111111")
	gifts := models.GetAllGifts()
	g.Data["json"] = gifts
	g.ServeJSON()
}

// @Title Get
// @Description get gift by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Gift
// @Failure 403 :id is empty
// @router /:id [get]

func (g * GiftController) Get() {
	id, e := g.GetInt("id", 0)
	if e != nil {
		g.Data["json"] = e.Error()
	} else {
		gift := models.GetGift(id)
		g.Data["json"] = gift
	}

	g.ServeJSON()
}