package main

import (
	"fmt"

	"github.com/msack22/go/variables"
)

func main() {
	variables.RestoVariables()

	estado, text := variables.ConviertoaTexto(10)
	fmt.Println(estado)
	fmt.Println(text)
}
