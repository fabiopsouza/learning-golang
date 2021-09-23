package main

import (
	"fmt"
	"git/learning-golang/2-banco/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64)
}

// --
// PARTE 1

// func main() {
// 	// contaDoGuilherme := ContaCorrente{
// 	// 	titular:       "Guilherme",
// 	// 	numeroAgencia: 589,
// 	// 	numeroConta:   123456,
// 	// 	saldo:         125.50}

// 	contaDoGuilherme := ContaCorrente{"Guilherme", 589, 123456, 125.50}
// 	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
// 	contaDaBruna2 := ContaCorrente{"Bruna", 222, 111222, 200}

// 	fmt.Println(contaDaBruna == contaDaBruna2) // true -> Compara valores, não referência (equivalente ao equals do Java)

// 	fmt.Println(contaDoGuilherme, contaDaBruna)

// 	var contaDaCris *ContaCorrente
// 	contaDaCris = new(ContaCorrente) // Só pode atribuir new a ponteiros
// 	contaDaCris.titular = "Cris"

// 	var contaDaCris2 *ContaCorrente
// 	contaDaCris2 = new(ContaCorrente)
// 	contaDaCris2.titular = "Cris"

// 	fmt.Println(contaDaCris, contaDaCris2)

// 	fmt.Println(contaDaCris == contaDaCris2)   // false -> Apesar dos mesmos valores, está comparando os ponteiros
// 	fmt.Println(*contaDaCris == *contaDaCris2) // true -> *nome retorna o valor do ponteiro, portanto os valores são iguais
// }

// --
// PARTE 2

// func main() {
// 	contaDaSilvia := contas.ContaCorrente{"Silvia", 0, 0, 500}
// 	contaDoPedro := contas.ContaCorrente{"Pedro", 0, 0, 500}

// 	// contaDaSilvia.Sacar(300)
// 	// contaDaSilvia.Depositar(100)
// 	// contaDaSilvia.Depositar(200)

// 	contaDaSilvia.Transferir(50, &contaDoPedro)

// 	fmt.Println(contaDaSilvia, contaDoPedro)
// }

// --
// PARTE 3

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())
}
