package routers

import (
	"dddddd/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//router: 路由
	beego.Router("/", &controllers.RegisterController{})
	//用户注册接口
    beego.Router("/register", &controllers.MainController{})
	//用户登录的接口
	beego.Router("/login",&controllers.LoginController{})
	//请求直接登录的页面
	beego.Router("/login.html", &controllers.LoginController{})
	//用户上传文件的功能
	beego.Router("/upload",&controllers.UploadFileController{})
	//在认证数据列表页面，点击新增认证按钮，跳转"新增页面"
	beego.Router("/upload_file.html",&controllers.UploadFileController{})
}
