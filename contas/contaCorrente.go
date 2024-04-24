package contas

import (
	"github.com/cassioglay/banco/clientes"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	podeSadar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSadar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	}

	return "Erro, você não tem limite para essa operação!"
}

func (c *ContaCorrente) Depositar(valorDoDepositos float64) (string, float64) {

	if valorDoDepositos > 0 {
		c.saldo += valorDoDepositos
		return "Deposito realizado com sucesso!", c.saldo
	}

	return "Erro, você não pode depositar esse valor!", c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorDaTransferencia <= c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}

	return false
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
