package clases

import "time"

// Clase contiene la información de una sesión de surf.
type Clase struct {
    ID              int
    Nombre          string
    Fecha           time.Time
    NivelRequerido  string
    InstructorID    int
    CupoMaximo      int
}
