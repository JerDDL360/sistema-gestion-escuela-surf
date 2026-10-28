package instructores

import "fmt"

// Instructor guarda los datos de cada instructor de surf.
type Instructor struct {
	ID           int
	Nombre       string
	Telefono     string
	Especialidad string
}

// Esta función revisa los datos básicos del instructor.
func ValidarInstructor(i Instructor) error {
	if i.ID <= 0 {
		return fmt.Errorf("el ID debe ser mayor que cero")
	}

	if i.Nombre == "" {
		return fmt.Errorf("el nombre no puede estar vacío")
	}

	if i.Especialidad == "" {
		return fmt.Errorf("la especialidad no puede estar vacía")
	}

	return nil
}
