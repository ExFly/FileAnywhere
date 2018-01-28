package controllers

import (
	"net"
	"github.com/astaxie/beego"
	"github.com/skip2/go-qrcode"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "????"
	c.Data["Email"] = "????@gmail.com"
	c.TplName = "index.tpl"
}

type FileQRController struct {
	beego.Controller
}

func (this *FileQRController) Get() {
	filename := this.Ctx.Input.Param(":filename")
	beego.Debug(filename)
	this.Data["filename"] = filename
	
	filename_pre := "qr"
	qr_file_name := filename_pre+".png"
	file_path := "static/qr/"+qr_file_name

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		beego.Debug(err)
	}

	var ip string
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				beego.Debug(ipnet.IP.String())

			}
		}
	}

	file_url := "http://"+ ip +":8080/download/"+filename
	err = qrcode.WriteFile(file_url, qrcode.Medium, 256, file_path)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["file_url"] = file_url
		this.Data["file_qr"] = "/"+file_path
	}

	this.TplName = "file.tpl"
}
