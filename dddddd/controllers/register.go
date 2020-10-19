package controllers

import (
	"dddddd/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post(){
	//解析用户端提交的请求数据
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("抱歉，数据解析失败,请重试!")
		return
	}

	//2、将解析到的数据保存到数据库中
	_, err = user.AddUser()
	if err != nil {
		r.Ctx.WriteString("抱歉，用户注册失败，请重试!")
		return
	}

	//将处理结果返回给客户端浏览器
	// 如果成功，跳转登录页面
	//tpl: template
	r.TplName =  "login.html"
	//如果失败，则提示错误信息
}
