package centro 

import "sync"

// Centro representa el centro de fisioterapia con sus recursos y colas
type Centro struct {
	// Canales para representar a los empleados como semáforos
	FisioDisponible chan struct{}
	MasajistaDisponible chan struct{}


	// Canal para el vestuario, un recurso único.
	Vestuario chan struct{}

	// WaitGroup para sincronizar 
	Wg sync.WaitGroup
}

// Instanciamos Centro
func NewCentro() *Centro {
	c := &Centro{
		// Buffered channels de tamaño 1, un recurso libre
		FisioDisponible:			make(chan struct{}, 1),
		MasajistaDisponible:	make(chan struct{}, 1),
		Vestuario: 						make(chan struct{}, 1),
	}

	// Inicializamos los Canales
	c.FisioDisponible <- struct{}{}
	c.MasajistaDisponible <- struct{}{}

	return c
}
