package main

import (
	"fmt"
	"time"
)

func main() {
	// Criando um canal
	channel := make(chan string)

	// Iniciando uma goroutine para produzir dados no canal
	go func() {
		for i := 1; ; i++ {
			data := fmt.Sprintf("Mensagem %d", i)
			channel <- data
			time.Sleep(time.Second)
		}
	}()

	// Loop infinito para monitorar o canal
	for {
		select {
		case msg := <-channel:
			fmt.Println("Recebeu:", msg)
			// Faça o que quiser com a mensagem recebida
		}
	}
}
