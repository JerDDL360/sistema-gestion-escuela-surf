package alumnos

import "fmt"

// Alumno guarda los datos de cada estudiante.
type Alumno struct {
	ID       int
	Nombre   string
	Edad     int
	Telefono string
	Nivel    string
}

// Esta función revisa los datos antes de registrar al alumno.
func ValidarAlumno(a Alumno) error {
	if a.ID <= 0 {
		return fmt.Errorf("el ID debe ser mayor que cero")
	}
	if a.Nombre == "" {
		return fmt.Errorf("el nombre no puede estar vacío")
	}
	if a.Edad <= 0 {
		return fmt.Errorf("la edad debe ser mayor que cero")
	}

	switch a.Nivel {
	case "Principiante", "Intermedio", "Avanzado":
		return nil
	default:
		return fmt.Errorf("el nivel debe ser Principiante, Intermedio o Avanzado")
	}
}
