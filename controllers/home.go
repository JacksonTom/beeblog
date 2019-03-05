package controllers

import (
	"github.com/JacksonTom/beeblog/models"
	"github.com/astaxie/beego"
	"unicode/utf8"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	//this.Ctx.Output.Header("Cache-Control","max-age=0")
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics,err := models.GetAllTopics(this.Input().Get("cate"),this.Input().Get("label"),true)
	if err !=nil{
		beego.Error(err)
	}
	for i:=0;i<len(topics);i++{
		if utf8.RuneCountInString(topics[i].Content)>40{
			topics[i].Content=string([]rune(topics[i].Content)[0:40])+"..."
		}
	}
	this.Data["Topics"] = topics

	categories,err := models.GetAllCategories()
	if err!=nil{
		beego.Error(err)
	}
	this.Data["Categories"] = categories

}
