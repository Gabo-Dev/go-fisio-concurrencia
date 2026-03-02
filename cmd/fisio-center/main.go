package main
import (
	"fmt"
	"math/rand"
	"time"
	"github.com/Gabo-Dev/go-fisio-concurrencia/internal/centro"
	"github.com/Gabo-Dev/go-fisio-concurrencia/internal/cliente"
)
func main() {
	fmt.Println("Iniciando simulación del centro de fisioterapia...")
	miCentro := centro.NewCentro()
	numClientesTotales := 10 // Como pide el enunciado
	

	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Generando %d clientes aleatorios...\n\n", numClientesTotales)
	
	for i := 1; i <= numClientesTotales; i++ {
		// Definimos el tipo de cliente aleatoriamente
		tipo := cliente.Masaje
		if rand.Intn(100) < 40 { // 40% de probabilidad para Rehabilitacion
			tipo = cliente.Rehabilitacion
		}
		// Creamos el cliente
		nuevoCliente := &cliente.Cliente{
			ID:   i,
			Tipo: tipo,
            Assigned: make(chan string, 1), // Inicializamos el canal de asignación del cliente
		}
		// Le decimos al WaitGroup que estamos añadiendo una nueva goroutine.
		miCentro.Wg.Add(1)
		// Lanzamos la goroutine del cliente
		go nuevoCliente.Vivir(miCentro)
		// Intervalo de llegada aleatorio entre 1 y 2 segundos
		intervaloLlegada := time.Duration(rand.Intn(2)+1) * time.Second
		time.Sleep(intervaloLlegada)
	}
	fmt.Println("\n...Todos los clientes han sido generados. Esperando que todas las sesiones terminen...")
	// Esperamos a que todas las goroutines de clientes finalicen.
	miCentro.Wg.Wait()
	// Notificamos al dispatcher que debe cerrar para una salida limpia.
	close(miCentro.Done)
	time.Sleep(50 * time.Millisecond) // Damos un pequeño tiempo para que el dispatcher procese el cierre
	fmt.Println("\nTodos los clientes se han ido. El centro cierra por hoy.")
	fmt.Println("Simulación finalizada.")
}
