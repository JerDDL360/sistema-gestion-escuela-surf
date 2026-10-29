package main

import "fmt"

func main() {
    var opcion int

    // Aquí se muestra el menú para entrar a cada parte de la escuela.
    // Se repite para que no tengamos que abrir el programa cada vez.
    for {
        fmt.Println("\nSistema de Gestión para una Escuela de Surf")
        fmt.Println("1. Gestión de alumnos")
        fmt.Println("2. Gestión de instructores")
        fmt.Println("3. Gestión de clases")
        fmt.Println("4. Gestión de reservas")
        fmt.Println("5. Seguimiento deportivo")
        fmt.Println("6. Salir")
        fmt.Print("Seleccione una opción: ")

        // Guardamos el número que el usuario escribe.
        _, err := fmt.Scanln(&opcion)
        if err != nil {
            fmt.Println("Por favor, ingrese un número válido.")
            return
        }

        // Según el número que se elija, mostramos el módulo correspondiente.
        switch opcion {
        case 1:
            fmt.Println("Módulo de alumnos seleccionado.")
        case 2:
            fmt.Println("Módulo de instructores seleccionado.")
        case 3:
            fmt.Println("Módulo de clases seleccionado.")
        case 4:
            fmt.Println("Módulo de reservas seleccionado.")
        case 5:
            fmt.Println("Módulo de seguimiento deportivo seleccionado.")
        case 6:
            fmt.Println("Saliendo del sistema...")
            return
        default:
            // Si se escribe un número que no está en el menú, avisamos.
            fmt.Println("Opción no válida. Intente nuevamente.")
        }
    }
}
