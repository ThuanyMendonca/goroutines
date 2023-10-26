package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Implementar forma de esperar uma função terminar a execução
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			DemorarParaExecutar(2 * time.Second)
		}()
	}

	wg.Wait()

	// E se eu registrar registradores a menos no wait group?
	// inicializa as 5 e finaliza só 3 (ex)

	// E se eu registrar valores a mais no wait group?
	// irá executar as 5 e dar um fatal error deadlock que mata a aplicação
}

func DemorarParaExecutar(t time.Duration) {
	fmt.Println("Começando a execução do método")

	time.Sleep(t)

	fmt.Println("Finalizando a execução do método")
}
