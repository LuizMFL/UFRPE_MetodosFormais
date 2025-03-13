/*
O problema do Produtor-Consumidor é um exemplo clássico de concorrência no qual alguns processos devem produzir itens e outros devem consumir os itens produzidos.
Em anexo existe uma especificação simplificada desta problema com um processo Produtor e dois processos Consumidores.
Também segue em anexo um código em Go com uma possível implementação incompleta desta especificação CSP.
Você deve ler os comentários no código e completá-lo de forma a ficar consistente com a especificação CSP.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// declare uma variável chamada ch que é um canal de inteiros
var ch chan int

func produtor() {
	rand.Seed(time.Now().UnixNano())
	for {
		time.Sleep(1000 * time.Millisecond)
		v := rand.Intn(5)
		fmt.Printf("produz.%v,\n", v)
		// use ch para enviar o valor produzido no canal ch
		ch <- v
	}
}

func consumidor(id int) {
	for {
		time.Sleep(1000 * time.Millisecond)
		// use ch para receber o valor enviado pelo produtor
		// guarde o valor na variável v
		v := <-ch
		fmt.Printf("consome.%v.%v,\n", id, v)
	}
}

func main() {

	// use o comando make para inicializar o canal aqui
	ch = make(chan int, 5)
	go produtor()

	for i := 0; i < 2; i++ {
		go consumidor(i)
	}

	// use time.Sleep para determinar o tempo que main roda
	time.Sleep(10 * time.Second)
}
