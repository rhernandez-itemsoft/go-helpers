package main

import (
	"fmt"

	"github.com/rhernandez-itemsoft/go-helpers/files/toml"
)

func main() {
	fmt.Print("test")
	toml.Get("config/app.tml")
}
