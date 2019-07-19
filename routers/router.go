package routers

import (
	"repetition/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/index", &controllers.MainController{}, "get:ShowIndex")
	beego.Router("/addArticle", &controllers.MainController{}, "get:ShowAdd;post:HandleAdd")
	beego.Router("/content", &controllers.MainController{}, "get:ShowContent;post:HandleContent")
	beego.Router("/update", &controllers.MainController{}, "get:ShowUpdate;post:HandleUpdate")
}
