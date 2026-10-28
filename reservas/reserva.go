package reservas

import "fmt"

// Reserva guarda los datos de la inscripción a una clase.
type Reserva struct {
	ID       int
	AlumnoID int
	ClaseID  int
	Estado   string
}

// Esta función revisa los datos básicos de la reserva.
func ValidarReserva(r Reserva) error {
	if r.ID <= 0 {
		return fmt.Errorf("el ID debe ser mayor que cero")
	}

	if r.AlumnoID <= 0 {
		return fmt.Errorf("se debe indicar un alumno válido")
	}

	if r.ClaseID <= 0 {
		return fmt.Errorf("se debe indicar una clase válida")
	}

	switch r.Estado {
	case "Confirmada", "Cancelada":
		return nil
	default:
		return fmt.Errorf("el estado debe ser Confirmada o Cancelada")
	}
}
