package main

import (
	"github.com/Noah-Wilderom/cfx-go"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	cfx.Print("cfxgo lib loaded")
	js.Global.Get("console").Call("log", "gopherjs load")
}