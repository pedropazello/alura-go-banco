package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoGuilherme := ContaCorrente{titular: "Guilherme", numeroAgencia: 589, numeroConta: 12345, saldo: 125.5}
	fmt.Println(contaDoGuilherme)

	contaDaCris := new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	fmt.Println(*contaDaCris)
}
