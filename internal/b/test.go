package b

import (
	"fmt"
	"project/internal/ctx"
)

type B struct{}

func Call(ctx ctx.Context) {
	ctx.A.PrintText()
}

func (_ B) PrintText() {
	fmt.Println("Teste from b")
}
