package fundamentos

import (
	"fmt"
	"time"
)

func declaracao() {
	var nome string = "Luis Munaira Junior"
	email := "munaira@gmail.com"
	datacriacao := time.Now()
	const ativo = true
	fmt.Println(nome, email, datacriacao)

	const (
		PI              = 3.14
		Juros           = 0.05
		parcelasMaximas = 12
	)
	fmt.Println(PI, Juros, parcelasMaximas)

}
