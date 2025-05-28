package fundamentos

import "fmt"

var n1 int = 10
var n2 int = 20

func artimetrico() {
	fmt.Println("Numero 1:", n1, "Numero 2:", n2)
	fmt.Println("Operadores Aritméticos")
	fmt.Println("Soma:", n1+n2)
	fmt.Println("Subtração:", n1-n2)
	fmt.Println("Multiplicação:", n1*n2)
	fmt.Println("Divisão:", n1/n2)
	fmt.Println("Módulo:", n1%n2)
	fmt.Println("Incremento:", n1+1)
	fmt.Println("Decremento:", n2-1)
	fmt.Println("Exponenciação:", n1*n1)
	fmt.Println("Negação:", -n1)
	fmt.Println("Positivo:", +n2)
	fmt.Println("Concatenação:", "Olá, "+"Mundo!")
	fmt.Println("Concatenação com número:", "Número: "+fmt.Sprint(n1))
	fmt.Println("Concatenação com float:", "Número: "+fmt.Sprintf("%.2f", float64(n1)))
	fmt.Println("Concatenação com booleano:", "Ativo: "+fmt.Sprint(true))
	fmt.Println("Concatenação com string:", "Nome: "+"Luis Munaira Junior")
	fmt.Println("Concatenação com string e número:", "Nome: "+"Luis Munaira Junior"+" - Número: "+fmt.Sprint(n1))
	fmt.Println("Concatenação com string e float:", "Nome: "+"Luis Munaira Junior"+" - Número: "+fmt.Sprintf("%.2f", float64(n1)))
	fmt.Println("Concatenação com string e booleano:", "Nome: "+"Luis Munaira Junior"+" - Ativo: "+fmt.Sprint(true))

}

func relacionais() {
	fmt.Println("Operadores Relacionais")
	fmt.Println("Igualdade:", n1 == n2)
	fmt.Println("Desigualdade:", n1 != n2)
	fmt.Println("Maior que:", n1 > n2)
	fmt.Println("Menor que:", n1 < n2)
	fmt.Println("Maior ou igual a:", n1 >= n2)
	fmt.Println("Menor ou igual a:", n1 <= n2)

}

func operadoreslogicos() {
	fmt.Println("Operadores Lógicos")
	fmt.Println("E lógico (AND):", n1 > 5 && n2 < 30)
	fmt.Println("Ou lógico (OR):", n1 > 15 || n2 < 25)
	fmt.Println("Não lógico (NOT):", !(n1 == n2))
	fmt.Println("XOR lógico:", (n1 > 5) != (n2 < 30)) // XOR não é um operador direto em Go, mas pode ser simulado
}
