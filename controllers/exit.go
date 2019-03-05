package controllers

import (
	"github.com/astaxie/beego"
)

type ExitController struct{
	beego.Controller
}

func (this *ExitController) Get(){

	//this.Ctx.Output.Header("Cache-Control","max-age=0")

	beego.Info("exit")
	this.Ctx.SetCookie("uname", "", -1, "/")
	this.Ctx.SetCookie("pwd", "", -1, "/")
	this.Redirect("/",302)


}