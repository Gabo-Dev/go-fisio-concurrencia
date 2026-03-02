package centro
import (
	"sync"
	"time"
	"github.com/Gabo-Dev/go-fisio-concurrencia/internal/cliente"
)
type Centro struct {
	FisioDisponible     chan struct{}
	MasajistaDisponible chan struct{}
	ColaRehabilitacion  chan *cliente.Cliente
	ColaMasaje          chan *cliente.Cliente
	Wg                  sync.WaitGroup
	Done                chan struct{}
}
func NewCentro() *Centro {
	c := &Centro{
		FisioDisponible:     make(chan struct{}, 1),
		MasajistaDisponible: make(chan struct{}, 1),
		ColaRehabilitacion:  make(chan *cliente.Cliente),
		ColaMasaje:          make(chan *cliente.Cliente),
		Done:                make(chan struct{}),
	}
	c.FisioDisponible <- struct{}{}
	c.MasajistaDisponible <- struct{}{}
	go c.runDispatcher()
	return c
}
func (c *Centro) WgDone() {
	c.Wg.Done()
}
func (c *Centro) SendToMasajeQueue(cli *cliente.Cliente) {
	c.ColaMasaje <- cli
}
func (c *Centro) SendToRehabilitacionQueue(cli *cliente.Cliente) {
	c.ColaRehabilitacion <- cli
}
func (c *Centro) ReleaseFisio() {
	c.FisioDisponible <- struct{}{}
}
func (c *Centro) ReleaseMasajista() {
	c.MasajistaDisponible <- struct{}{}
}
func (c *Centro) runDispatcher() {
	for {
		select {
		case <-c.Done:
			return
		case cli := <-c.ColaRehabilitacion:
			select {
			case <-c.FisioDisponible:
				cli.Assigned <- "FISIOTERAPEUTA"
			default:
				go func() { c.ColaRehabilitacion <- cli }()
			}
		default:
			select {
			case cli := <-c.ColaMasaje:
				select {
				case <-c.FisioDisponible:
					cli.Assigned <- "FISIOTERAPEUTA"
				case <-c.MasajistaDisponible:
					cli.Assigned <- "MASAJISTA"
				default:
					go func() { c.ColaMasaje <- cli }()
				}
			default:
				// Evita CPU spin cuando no hay clientes en cola.
				time.Sleep(50 * time.Millisecond)
			}
		}
	}
}

