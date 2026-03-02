package main

import (
	"fmt"
	"github.com/Gabo-Dev/go-fisio-concurrencia/internal/centro"
	"github.com/Gabo-Dev/go-fisio-concurrencia/internal/cliente"
)

func main() {
	fmt.Println("Iniciando simulación el centro de fisioterapia...")
	miCentro := centro.NewCentro()

	numClientes := 3
	fmt.Printf("Lanzando %d clientes de prueba...\n\n", numClientes)

	for i := 1; i <= numClientes; i++ {
		miCentro.Wg.Add(1)
		nuevoCliente := &cliente.Cliente{
			ID:   i,
			Tipo: cliente.Masaje,
		}
		go nuevoCliente.Vivir(miCentro)
	}
	fmt.Println("\n...Todos los clientes en el centro. Esperando que terminen sus sesiones...")
	

	miCentro.Wg.Wait()
	fmt.Println("\nTodos los clientes se han ido. El centro cierra por hoy.")

 }
