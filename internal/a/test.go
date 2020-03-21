package a

import (
	"fmt"
	"project/internal/ctx"
)

type A struct{}

func Call(ctx ctx.Context) {
	ctx.B.PrintText()
}

func (_ A) PrintText() {
	fmt.Println("Teste from a")
}
