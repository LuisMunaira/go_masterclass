package utils

import (
	"fmt"
	"strings"
)

func Executar(secao string, funcoes ...func()) {
	for i, funcao := range funcoes {
		fmt.Printf("\n\n >>> %s -Exercicio %d <<< \n\n", strings.ToUpper(secao), i+1)
		funcao()
	}
}
