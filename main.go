package main

import (
	_ "repetition/models"
	_ "repetition/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
