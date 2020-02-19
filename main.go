package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/rhernandez-itemsoft/go-helpers/files/toml"
	"github.com/rhernandez-itemsoft/go-helpers/iresponse"
	"github.com/rhernandez-itemsoft/go-helpers/iresponse/irespmdl"
)

// I18nModel Definición para el config i18n
type I18nModel struct {
	Default      string
	URLParameter string
}

var _iresponse *iresponse.Definition

//SERVER servidor
const SERVER = "localhost:8082"

func main() {
	app := iris.New()

	fmt.Println(iris.Version)

	var i18Conf *I18nModel
	toml.Get("./config/i18n.tml", &i18Conf)
	fmt.Println("Test I18N ------------------------------")
	fmt.Println(i18Conf.Default)

	fmt.Println("Test iResponse ------------------------------")
	app.Handle("GET", "/", func(ctx iris.Context) {
		_iresponse = iresponse.New(ctx)
		var response irespmdl.Response = irespmdl.NewResponse()
		response.Messages = append(response.Messages, "Probando el http response")

		_iresponse.JSON(response)
	})
	app.Run(iris.Addr(SERVER))
}
