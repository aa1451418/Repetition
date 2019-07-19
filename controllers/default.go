package controllers

import (
	"repetition/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "register.html"
}

func (c *MainController) Post() {
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	if userName == "" || pwd == "" {
		beego.Info("data no by kong")
		c.Redirect("/register", 302)
		return
	}

	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Pwd = pwd
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("Please input data")
		c.Redirect("/register", 302)
		return
	}
	c.Redirect("/login", 302)
}

func (c *MainController) ShowLogin() {
	c.TplName = "login.html"
}

func (c *MainController) HandleLogin() {
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	if userName == "" || pwd == "" {
		beego.Info("user or password no by vacancy")
		c.Redirect("/login", 302)
		return
	}

	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	// user.Pwd = pwd

	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("data no by vacancy")
		c.Redirect("/login", 302)
		return
	}
	c.Redirect("/index", 302)
}

func (c *MainController) ShowIndex() {
	c.TplName = "index.html"
}

func (c *MainController) ShowAdd() {
	c.TplName = "add.html"
}

func (c *MainController) HandleAdd() {
	articleName := c.GetString("articleName")
	content := c.GetString("content")

	if articleName == "" || content == "" {
		beego.Info("不能插入空内容")
		return
	}

	o := orm.NewOrm()
	article := models.Article{}
	article.ArtiName = articleName
	article.Content = content
	_, err := o.Insert(&article)
	if err != nil {
		beego.Info("数据不错误")
		return
	}
	c.Redirect("/index", 302)
}

func (c *MainController) ShowContent() {
	c.TplName = "content.html"
}

func (c *MainController) HandleContent() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("get id erro", err)
		return
	}

	/* query */
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	err = o.Read(&arti)
	if err != nil {
		beego.Info("query erro", err)
	}

	c.TplName = "content.html"

}

/* 显示编辑页面 */
func (c *MainController) ShowUpdate() {
	c.TplName = "update.html"
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("Query Error", err)
		return
	}

	/* Query data  get data */
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	err = o.Read(&arti)
	if err != nil {
		beego.Info("Query erro")
		return
	}

}

func (c *MainController) HandleUpdate() {
	id, _ := c.GetInt("id")
	articleName := c.GetString("articleName")
	content := c.GetString("content")
	_, err := c.GetFile("uploadname")
	// var filename string
	if err != nil {
		beego.Info("数据错误")
		return
	}

	o := orm.NewOrm()
	arti := models.Article{Id: id}
	err = o.Read(&arti)
	if err != nil {
		beego.Info("Data Query Error")
	}

	arti.ArtiName = articleName
	arti.Content = content
	_, err = o.Update(&arti, "Article", "content")
	if err != nil {
		beego.Info("Please Correct data")
		return
	}
}
