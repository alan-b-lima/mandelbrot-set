//go:build js

package jsconsole

import (
	"fmt"
	"syscall/js"
)

type JSConsole struct {
	js.Value
}

type errorJSConsole struct {
	js.Value
}

var console = JSConsole{js.Global().Get("console")}

func Console() JSConsole {
	return console
}

func (c *JSConsole) Write(p []byte) (n int, err error) {
	str := string(p)
	c.Call("log", str)

	return len(p), nil
}

func (ec *errorJSConsole) Write(p []byte) (n int, err error) {
	str := string(p)
	ec.Call("error", str)

	return len(p), nil
}

func Log(a ...any) {
	fmt.Fprintln(&console, a...)
}

func Error(a ...any) {
	fmt.Fprintln(&errorJSConsole{console.Value}, a...)
}
