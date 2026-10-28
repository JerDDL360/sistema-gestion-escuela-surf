package reservas

// Reserva contiene los datos de una inscripción a una clase.
type Reserva struct {
    ID         int
    AlumnoID   int
    ClaseID    int
    Estado     string
}
