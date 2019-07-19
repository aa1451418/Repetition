# Repetition
repetition learn

![](E:\GOPATH\src\repetition\images\1563538273.jpg)

路由

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

