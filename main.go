package main

type verificaConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificaConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {

	/*
		cliente1 := clientes.Titular{Nome: "Bruno", CPF: "123.111.123.12", Profissao: "Desenvolvedor"}
		   	contaCliente1 := contas.ContaCorrente{Titular: cliente1, NumeroAgencia: 123, NumeroConta: 123456}

		   	fmt.Println(contaCliente1)

		   	contaCliente1.Depositar(100)

		   	fmt.Println(contaCliente1)
		   	fmt.Println(contaCliente1.ObterSaldo())


		contaCliente2 := contas.ContaPoupanca{}
		   	contaCliente2.Depositar(100)
		   	PagarBoleto(&contaCliente2, 60)

		   	fmt.Println(contaCliente2)
		   	fmt.Print(contaCliente2.ObterSaldo())
	*/

}
