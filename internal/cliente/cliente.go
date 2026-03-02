
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

	time.Sleep(1 * time.Second) // Simulamos que está un rato en el centro
	fmt.Printf("<< Cliente %d se ha ido.\n", cli.ID)
}
