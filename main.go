package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hola Mundo desde Go en Railway!")
}

func main() {
    http.HandleFunc("/", handler)
    
    // Obtiene el puerto de la variable de entorno, o usa el puerto 8080 por defecto
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Printf("Server running on port %s\n", port)
    http.ListenAndServe(":" + port, nil)
}
