package contas

import (
	"fmt"
	"git/learning-golang/2-banco/clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
	}

	fmt.Println(c.saldo)
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) {
	podeSacar := valorDoDeposito > 0
	if podeSacar {
		c.saldo += valorDoDeposito
	}

	fmt.Println(c.saldo)
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
