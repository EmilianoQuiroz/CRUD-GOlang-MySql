package main

import (
	"net/http" 		// Permite construir servidores HTTP 
	"log" 	   		// Nos permite crear diferentes tipos de loggers, usando el método New
	//"fmt" 	   		// Formato de entrada y salida de datos
	"text/template" // Para trabajar con templates
)
// Llamado a nuestro template
var plantillas = template.Must(template.ParseGlob("plantillas/*"))
func main() {
	//Solicitud para acceder a la funcion inicio
	http.HandleFunc("/", Inicio)
	// Log que indica por consola que el servidor esta corriendo
	log.Println("Servidor corriendo...")
	// Indicamos el servidor en el que estara corriendo la aplicacion
	http.ListenAndServe(":8080", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w,"Hola Go")
	// Accedemos al contenido de la plantilla inicio
	plantillas.ExecuteTemplate(w, "inicio", nil)
}