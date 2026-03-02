
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

	fmt.Printf(" ...Cliente  %d está en la cola, esperando un masajista...\n",cli.ID)
	<-c.MasajistaDisponible
  fmt.Printf("Masajista: Atendiendo al cliente %d\n", cli.ID)

	time.Sleep(2 * time.Second) // Simulamos que está un rato en el centro
	
	// liberamos canal
	c.MasajistaDisponible <- struct{}{}
	fmt.Printf("   <-- Cliente %d TERMINÓ su masaje y LIBERA al masajista.\n", cli.ID)
	fmt.Printf("<< Cliente %d se ha ido del centro.\n", cli.ID)
}
