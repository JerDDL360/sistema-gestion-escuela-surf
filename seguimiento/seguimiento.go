package seguimiento

// Evaluacion registra el progreso de un alumno en una clase.
type Evaluacion struct {
    ID             int
    AlumnoID       int
    ClaseID        int
    Equilibrio     int
    Remada         int
    ControlTabla   int
    Observaciones  string
}
