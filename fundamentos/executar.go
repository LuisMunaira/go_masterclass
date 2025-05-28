package fundamentos

import (
	"fmt"

	"github.com/LuisMunaira/go_masterclass.git/utils"
)

func Executar() {
	fmt.Println("Executando o pacote fundamentos")
	utils.Executar("fundamentos", primeiro, tipos,
		declaracao,
		artimetrico,
		relacionais,
		operadoreslogicos,
	)
	fmt.Println("Fim do pacote fundamentos")

	//primeiro()
	//tipos(  )
}
