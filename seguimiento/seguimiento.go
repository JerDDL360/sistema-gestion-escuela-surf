package seguimiento

import "fmt"

// Evaluacion guarda los resultados de un alumno durante una clase.
type Evaluacion struct {
	ID            int
	AlumnoID      int
	ClaseID       int
	Equilibrio    int
	Remada        int
	ControlTabla  int
	Observaciones string
}

// Esta función revisa las puntuaciones de la evaluación.
func ValidarEvaluacion(e Evaluacion) error {
	if e.ID <= 0 {
		return fmt.Errorf("el ID debe ser mayor que cero")
	}

	if e.AlumnoID <= 0 || e.ClaseID <= 0 {
		return fmt.Errorf("se debe indicar un alumno y una clase válidos")
	}

	if e.Equilibrio < 1 || e.Equilibrio > 5 {
		return fmt.Errorf("el equilibrio debe tener una puntuación del 1 al 5")
	}

	if e.Remada < 1 || e.Remada > 5 {
		return fmt.Errorf("la remada debe tener una puntuación del 1 al 5")
	}

	if e.ControlTabla < 1 || e.ControlTabla > 5 {
		return fmt.Errorf("el control de la tabla debe tener una puntuación del 1 al 5")
	}

	return nil
}
