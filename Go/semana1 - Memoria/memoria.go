/*
O arquivo golang em anexo é uma versão inicial descrita na atividade de paralelismo em CSP, na qual tínhamos 3 componentes e uma memória.
Atualmente o arquivo só tem um componente e uma memória.

Complete o arquivo golang em anexo com os dois componentes faltando.
Faça alterações no programa para verificar os diferentes comportamentos que podem ser gerados.

Comp 1 - ele pode ler a memória e guardar em um parâmetro do processo, ou caso este parâmetro seja maior que cinco ele deve imprimir o evento "esquerda", e caso ele seja menor ou igual a cinco deve imprimir o evento "direita".
Comp 2 - pode ler o valor da memória, e caso ele seja maior que cinco, ele decrementa em um o valor na memória, caso ele seja menor  ou igual a cinco ele incrementa em um o valor da memória.
Comp 3 - pode ler o valor da memória e caso ele seja menor ou igual a cinco, ele decrementa em um o valor da memória, e caso ele seja maior que cinco ele incrementa em um o valor da memória.
Memória - controla o acesso a um inteiro entre 0 e 10 permitindo a leitura ou escrita de somente um componente por vez.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func memo(read, write chan int) {
	x := 0
	for {
		select {
		case read <- x:
			fmt.Println("[M] Leu x =", x)
		case y := <-write:
			x = x + y
			fmt.Println("[M] Escreveu x =", x)
		}
	}
}
func comp1(read chan int) {
	value := 0
	for {
		select {
		case value = <-read:
			fmt.Println("[1] Leu da memoria o valor ", value)
		default:
			if value > 5 {
				fmt.Println("[1] Esquerda")
			} else {
				fmt.Println("[1] Direita")
			}
			n := rand.Intn(500)
			time.Sleep(time.Duration(n) * time.Millisecond)
		}
	}
}

func comp2(read, write chan int) {
	value := 0
	for {
		select {
		case value = <-read:
			fmt.Println("[2] Leu da memoria o valor ", value)
		default:
			if value > 5 {
				write <- (-1)
				value = value - 1
			} else {
				write <- 1
				value = value + 1
			}
			n := rand.Intn(500)
			time.Sleep(time.Duration(n) * time.Millisecond)
		}
	}
}

func comp3(read, write chan int) {
	value := 0
	for {
		select {
		case value = <-read:
			fmt.Println("[3] Leu da memoria o valor ", value)
		default:
			if value <= 5 {
				write <- (-1)
				value = value - 1
			} else {
				write <- 1
				value = value + 1
			}
			n := rand.Intn(500)
			time.Sleep(time.Duration(n) * time.Millisecond)
		}
	}
}

func main() {
	read := make(chan int)
	write := make(chan int)
	go comp1(read)
	go comp2(read, write)
	go comp3(read, write)
	go memo(read, write)
	time.Sleep(10 * time.Second)
}
