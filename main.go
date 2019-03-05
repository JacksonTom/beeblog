package main

import (
	"github.com/JacksonTom/beeblog/models"
	"github.com/astaxie/beego"
	"os"

	_ "github.com/JacksonTom/beeblog/routers"
)

func init() {
	models.RegisterDB()
}

func main() {
	os.Mkdir("attachment",os.ModePerm)
	beego.Run()
}