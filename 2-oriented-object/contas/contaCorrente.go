package contas

import (
	"fmt"
	"git/learning-golang/2-banco/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
	}

	fmt.Println(c.saldo)
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) {
	podeSacar := valorDoDeposito > 0
	if podeSacar {
		c.saldo += valorDoDeposito
	}

	fmt.Println(c.saldo)
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) {
	if valorDaTransferencia < c.saldo {
		c.Sacar(valorDaTransferencia)
		contaDestino.Depositar(valorDaTransferencia)
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
