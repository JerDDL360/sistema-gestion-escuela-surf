package clases

import (
	"fmt"
	"time"
)

// Clase guarda los datos de cada sesión de surf.
type Clase struct {
	ID             int
	Nombre         string
	Fecha          time.Time
	NivelRequerido string
	InstructorID   int
	CupoMaximo     int
}

// Esta función revisa algunos datos de la clase.
func ValidarClase(c Clase) error {
	if c.ID <= 0 {
		return fmt.Errorf("el ID debe ser mayor que cero")
	}

	if c.Nombre == "" {
		return fmt.Errorf("el nombre de la clase no puede estar vacío")
	}

	if c.InstructorID <= 0 {
		return fmt.Errorf("se debe indicar un instructor válido")
	}

	if c.CupoMaximo <= 0 {
		return fmt.Errorf("el cupo máximo debe ser mayor que cero")
	}

	switch c.NivelRequerido {
	case "Principiante", "Intermedio", "Avanzado":
		return nil
	default:
		return fmt.Errorf("el nivel debe ser Principiante, Intermedio o Avanzado")
	}
}
