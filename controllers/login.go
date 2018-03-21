package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Post()  {
	/*
	打印网页输入的账户密码
	c.Ctx.WriteString(fmt.Sprint(c.Input()))
	return
    */
    uname := c.Input().Get("uname")
    pwd := c.Input().Get("pwd")
    autologin := c.Input().Get("autologin") == "on"
    if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd {
    	maxAge := 0
    	if autologin{
    		maxAge =1<<31-1
		}
		c.Ctx.SetCookie("uname",uname,maxAge,"/")
		c.Ctx.SetCookie("pwd",pwd,maxAge,"/")
	}
	c.Redirect("/",301)
	return
}

func checklogin (ctx *context.Context) bool {
	ck,err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	pck,err := ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := pck.Value
	return beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd
}