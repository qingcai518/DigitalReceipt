package controllers

import (
	"digitalReceipt/models"
	"github.com/astaxie/beego"
	"encoding/json"
)

// Operations about Gifts
type ItemController struct {
	 beego.Controller
}

// @Title GetAll
// @Description get all Receipts
// @Success 200 {object} models.Receipt
// @router / [get]
func (controller *ItemController) GetAll(receiptId int) {
	items := models.GetAllItems(receiptId)
	controller.Data["json"] = items
	controller.ServeJSON()
}

// @Title CreateReceipt
// @Description create Receipt
// @Param	body		body 	models.Receipt	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (controller *ItemController) Post() {
	var item models.Item
	var err = json.Unmarshal(controller.Ctx.Input.RequestBody, &item)
	if err != nil {
		controller.Data["json"] = err.Error()
	} else {
		itemId, err := models.AddItem(item)
		if err != nil {
			controller.Data["json"] = err.Error()
		}
		println(itemId)
		controller.Data["json"] = map[string]int64{"ItemId": itemId}
	}
	controller.ServeJSON()
}
