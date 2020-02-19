package main

import (
	"fmt"

	"github.com/rhernandez-itemsoft/go-helpers/files/toml"
)

// I18nModel Definición para el config i18n
type I18nModel struct {
	Default      string
	URLParameter string
}

func main() {
	var response *I18nModel
	toml.Get("./config/i18n.tml", &response)

	fmt.Println("Test I18N ------------------------------")
	fmt.Println(response.Default)
	fmt.Println(response.URLParameter)

}
