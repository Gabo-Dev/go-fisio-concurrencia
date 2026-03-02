
package cliente
import (
	"fmt"
	"time"
	"github.com/Gabo-Dev/go-fisio-concurrencia/internal/centro"
)

type TipoCliente int

const (
	Masaje TipoCliente = iota
	Rehabilitacion
)

func (tc TipoCliente) String() string {
	return [...]string{"Masaje", "Rehabilitacion"}[tc]
}

type Cliente struct{
	ID int
	Tipo TipoCliente
}

func (cli *Cliente) Vivir(c *centro.Centro){
	defer c.Wg.Done()

	fmt.Printf(">> Cliente %d ha llegado para una sesión de %s.\n", cli.ID, cli.Tipo)

	fmt.Printf(" ...Cliente  %d está en la cola, esperando su turno para ser atendido por un empleado.\n",cli.ID)
	select {
		case <-c.MasajistaDisponible:
  		fmt.Printf("Masajista: Atendiendo al cliente %d\n", cli.ID)
			time.Sleep(2 * time.Second) // Simulamos que está un rato en el centro
			// liberamos canal del masajista
			c.MasajistaDisponible <- struct{}{}
			fmt.Printf("   <-- Cliente %d TERMINÓ su masaje y LIBERA al masajista.\n", cli.ID)
		case <-c.FisioDisponible:
			fmt.Printf("Fisio: Atendiendo al cliente %d\n", cli.ID)
			time.Sleep(2 * time.Second)
			// liberamos canal del Fisio
			c.FisioDisponible <- struct{}{}
			fmt.Printf("   <-- Cliente %d TERMINÓ su masaje y LIBERA al FISIOTERAPEUTA.\n", cli.ID)
	}
	fmt.Printf("<< Cliente %d se ha ido del centro.\n", cli.ID)
}
