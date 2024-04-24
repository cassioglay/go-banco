package contas

import (
	"github.com/cassioglay/banco/clientes"
)

type ContaPoupanca struct {
	Titular                               clientes.Titular
	NumeroAgencia, NumeroConta, Operaccao int
	saldo                                 float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {

	podeSadar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSadar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	}

	return "Erro, você não tem limite para essa operação!"
}

func (c *ContaPoupanca) Depositar(valorDoDepositos float64) (string, float64) {

	if valorDoDepositos > 0 {
		c.saldo += valorDoDepositos
		return "Deposito realizado com sucesso!", c.saldo
	}

	return "Erro, você não pode depositar esse valor!", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
