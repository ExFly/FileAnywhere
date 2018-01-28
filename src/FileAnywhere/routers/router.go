package routers

import (
    "FileAnywhere/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/file/:filename([\\w\\.]+)", &controllers.FileQRController{})
}
