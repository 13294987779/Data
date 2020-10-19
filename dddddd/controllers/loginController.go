package controllers

import (
	"dddddd/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController)Post()  {
    //解析客户端用户提交 的登入数据
	var user models.User
	err :=l.ParseForm(&user)
	if err !=nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("用户登入信息解析失败")
		return
	}

	//根据解析到的数据，执行数据库查询操作
u,err :=user.QueryUser()
	//判断数据库查询结果
if err !=nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("用户登入失败")
		return
	}

	//根据查询结果返回客户端相应的信息或者页面跳转
	l.Data["Phone"] =u.Phone//动态数据设置
	l.TplName ="home.html"//文件上传界面
}
