package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct{
	beego.Controller
}

func (this *LoginController) Get(){

	//this.Ctx.Output.Header("Cache-Control","max-age=0")

	this.TplName = "login.html"
}

func (this *LoginController) Post(){
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd{
		maxAge := 0
		if autoLogin{
			maxAge =31536000
		}
		this.Ctx.SetCookie("uname",uname,maxAge,"/")
		this.Ctx.SetCookie("pwd",pwd,maxAge,"/")
	}
	this.Redirect("/",302)
	return
}

func checkAccount(ctx *context.Context) bool{
	ck,err:=ctx.Request.Cookie("uname")
	if err!=nil{
		return false
	}
	uname :=ck.Value

	ck,err = ctx.Request.Cookie("pwd")
	if err !=nil{
		return false
	}
	pwd := ck.Value
	return beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd
}