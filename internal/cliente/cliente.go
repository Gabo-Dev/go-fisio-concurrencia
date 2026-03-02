package cliente
import (
	"fmt"
	"math/rand"
	"time"
)
type TipoCliente int
const (
	Masaje TipoCliente = iota
	Rehabilitacion
)
func (tc TipoCliente) String() string {
	return [...]string{"Masaje", "Rehabilitación"}[tc]
}
type Cliente struct {
	ID       int
	Tipo     TipoCliente
	Assigned chan string
}
type CenterService interface {
	WgDone()
	SendToMasajeQueue(cli *Cliente)
	SendToRehabilitacionQueue(cli *Cliente)
	ReleaseFisio()
	ReleaseMasajista()
}
func (cli *Cliente) Vivir(s CenterService) {
	defer s.WgDone()
	fmt.Printf(">> Cliente %d (%s) ha llegado.\n", cli.ID, cli.Tipo)
	if cli.Tipo == Masaje {
		s.SendToMasajeQueue(cli)
	} else {
		s.SendToRehabilitacionQueue(cli)
	}
	empleadoAsignado := <-cli.Assigned
	fmt.Printf("   --> Cliente %d (%s) fue asignado a %s.\n", cli.ID, cli.Tipo, empleadoAsignado)
	tiempoSesion := time.Duration(rand.Intn(3)+2) * time.Second
	fmt.Printf("      ... Cliente %d en sesión con %s por %v ...\n", cli.ID, empleadoAsignado, tiempoSesion)
	time.Sleep(tiempoSesion)
	if empleadoAsignado == "FISIOTERAPEUTA" {
		s.ReleaseFisio()
	} else {
		s.ReleaseMasajista()
	}
	fmt.Printf("   <-- Cliente %d liberó a %s.\n", cli.ID, empleadoAsignado)
	fmt.Printf("<< Cliente %d (%s) se ha ido del centro.\n", cli.ID, cli.Tipo)
}
