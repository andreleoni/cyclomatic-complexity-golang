package main

import (
	"project/internal/a"
	"project/internal/b"
	"project/internal/ctx"
)

func main() {
	app := ctx.Context{
		A: a.A{},
		B: b.B{},
	}

	a.Call(app)
	b.Call(app)
}
