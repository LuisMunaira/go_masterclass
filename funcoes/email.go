package funcoes

import "strings"

func testeEmail() {
	email := "munaira@gmail.com"
	usuario, dominio := partesDoemail(email)
	println("Usuário:", usuario)
	println("Domínio:", dominio)
}
func partesDoemail(email string) (usuario, dominio string) {
	partes := strings.Split(email, "@")
	usuario = partes[0]
	dominio = partes[1]
	return usuario, dominio

}
