package main

import (
	"net/http" 			// Permite construir servidores HTTP 
	"log" 	   			// Nos permite crear diferentes tipos de loggers, usando el método New
	//"fmt" 	   		// Formato de entrada y salida de datos
	"text/template" 	// Para trabajar con templates
	_"github.com/go-sql-driver/mysql" // Driver para la coneccion con la base de datos 
)
// Llamado a nuestro template
var plantillas = template.Must(template.ParseGlob("plantillas/*"))
func main() {
	// Solicitud para acceder a la funcion inicio
	http.HandleFunc("/", Inicio)
	// Solicitud para mostrar la plantilla crear
	http.HandleFunc("/crear", Crear)
	// Log que indica por consola que el servidor esta corriendo
	log.Println("Servidor corriendo...")
	// Indicamos el servidor en el que estara corriendo la aplicacion
	http.ListenAndServe(":8080", nil)
}

// Funcion para mostrar la plantilla de Inicio
func Inicio(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w,"Hola Go")
	// Accedemos al contenido de la plantilla inicio
	plantillas.ExecuteTemplate(w, "inicio", nil)
}

// Funcion para mostrar la plantilla de crear
func Crear(w http.ResponseWriter, r *http.Request){
	// Accedemos al contenido de la plantilla crear
	plantillas.ExecuteTemplate(w, "crear", nil)
}