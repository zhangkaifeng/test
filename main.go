package main

import (
	_ "project/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"project/models"
	"project/controllers"
	"time"
)
func init(){
	models.RegisterDB()
}
func main() {
	orm.Debug = true
	orm.DefaultTimeLoc = time.Local
	//自动建表
	orm.RunSyncdb("default",false,true)
	//注册beego路由
	beego.Router("/",&controllers.MainController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/fenlei",&controllers.Fenlei{})
	beego.Router("/topic",&controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
    //启动beego
	beego.Run()
}

