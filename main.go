package main

import (
	"fmt"

	"github.com/pedropazello/alura-go-banco/clientes"
	"github.com/pedropazello/alura-go-banco/contas"
)

func main() {
	titularGuilherme := clientes.Titular{}
	titularGuilherme.Nome = "Guilherme"
	titularGuilherme.CPF = "123456"
	titularGuilherme.Profissao = "Tecnico de eletronica"

	contaDoGuilherme := contas.ContaCorrente{Titular: titularGuilherme, NumeroAgencia: 589, NumeroConta: 12345}
	contaDoGuilherme.Depositar(500)

	PagarBoleto(&contaDoGuilherme, 500)

	fmt.Println(contaDoGuilherme)

	titular := clientes.Titular{}
	titular.Nome = "Cris"
	titular.CPF = "123456"
	titular.Profissao = "Tecnico de eletronica"

	contaDaCris := contas.ContaPoupanca{}
	contaDaCris.Titular = titular
	contaDaCris.Depositar(500)

	fmt.Println(contaDaCris)

	msg := contaDaCris.Sacar(100)

	fmt.Println(msg)
	fmt.Println(contaDaCris)

	msg2, saldo := contaDaCris.Depositar(100)

	fmt.Println(saldo)
	fmt.Println(msg2)
	fmt.Println(contaDaCris)

	contaDaCris.Transferir(100, &contaDoGuilherme)
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
