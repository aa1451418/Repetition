# Repetition
repetition learn

![](https://github.com/aa1451418/Repetition/blob/master/images/1563538273.jpg?raw=true)

beego路由

```go
func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/index", &controllers.MainController{}, "get:ShowIndex")
	beego.Router("/addArticle", &controllers.MainController{}, "get:ShowAdd;post:HandleAdd")
	beego.Router("/content", &controllers.MainController{}, "get:ShowContent;post:HandleContent")
	beego.Router("/update", &controllers.MainController{}, "get:ShowUpdate;post:HandleUpdate")
}

```



Post获取数据

![](https://github.com/aa1451418/Repetition/blob/master/images/20190719205031.png?raw=true)





















