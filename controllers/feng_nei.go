package controllers

import (
	"github.com/astaxie/beego"
	"project/models"
)

type Fenlei struct {
	beego.Controller
}

func (c *Fenlei) Get() {
	c.TplName = "fenlei.html"
	c.Data["Isfenlei"] = true

	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}

		err := models.Addfenlei(name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/fenlei", 301)
		return

	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
	}

    var err error
    c.Data["Fenlei"],err = models.GetAllfenlei()

    if err != nil {
    	beego.Error(err)
	}

}