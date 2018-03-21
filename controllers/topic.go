package controllers

import (
	"github.com/astaxie/beego"
	"project/models"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["IsLogin"] = checklogin(c.Ctx)
	c.Data["Iscont"] = true
	c.TplName = "topic.html"
	topics,err := models.GetAlltopic()
	if err != nil{
		beego.Error(err.Error)
	} else {
		c.Data["Topics"] = topics
	}
}

func (c *TopicController) Post() {
	if !checklogin(c.Ctx){
		c.Redirect("/login",302)
		return
	}

	title := c.Input().Get("title")
	content := c.Input().Get("content")

	var err error
	err = models.AddTopic(title,content)
	if err != nil{
		beego.Error(err)
	}
	c.Redirect("/topic",302)

}

func (c *TopicController) Add() {
	c.TplName = "topic_add.html"
}