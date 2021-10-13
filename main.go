package main

import (
	"fmt"

	"github.com/pedropazello/alura-go-banco/contas"
)

func main() {
	contaDoGuilherme := contas.ContaCorrente{Titular: "Guilherme", NumeroAgencia: 589, NumeroConta: 12345, Saldo: 125.5}
	fmt.Println(contaDoGuilherme)

	contaDaCris := contas.ContaCorrente{}
	contaDaCris.Titular = "Cris"
	contaDaCris.Saldo = 500

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
