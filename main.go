package main

import (
	"autoXreader/src"
	"autoXreader/src/selenium"
	"github.com/labstack/gommon/color"
)

func main() {
	color.Print(color.Cyan("Input UserName:"))
	userName := src.CmdLineInput()
	color.Print(color.Cyan("Input Password:"))
	userPass := src.CmdLineInput()
	selenium.Init(userName, userPass)
}
