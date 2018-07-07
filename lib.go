package main

import (
	"gopkg.in/russross/blackfriday.v2"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Get("document").Call("getElementById", "input").Call("addEventListener", "keyup", js.NewCallback(convert))
	<-c
}

func format(input string) string {
	return string(blackfriday.Run([]byte(input)))
}

func convert(_ []js.Value) {
	input := js.Global().Get("document").Call("getElementById", "input").Get("value").String()
	result := format(input)
	js.Global().Get("document").Call("getElementById", "output").Set("innerHTML", result)
}
