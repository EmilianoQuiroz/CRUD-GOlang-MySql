package main

import (
	"database/sql"		// Modulo necesario para trabajar con sql
	"net/http" 			// Permite construir servidores HTTP 
	"log" 	   			// Nos permite crear diferentes tipos de loggers, usando el método New
	//"fmt" 	   		// Formato de entrada y salida de datos
	"text/template" 	// Para trabajar con templates
	_"github.com/go-sql-driver/mysql" // Driver para la coneccion con la base de datos 
)

// Funcion para conectar con la base de datos
func conexionBD()(conexion *sql.DB){
	Driver:="mysql"
	Usuario:="root"
	Contrasenia:=""
	Nombre:="sistema"

	// Manejo de error en la DB
	conexion,err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(localhost:3307)/"+Nombre)
	
	if err != nil { // Si se produce un error
		//Ejecutamos el siguiente panic
		panic(err.Error())
	}
	return conexion
}

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

	// Prueba de conexion con la base de datos
	conexionEstablecida:= conexionBD()
	insertarRegistros,err:= conexionEstablecida.Prepare("INSERT INTO empleados(nombre,correo) VALUES('Santiago','correo@gmail.com')")

	// Para ejecutar la variable insertarRegistros, primero hay que asegurarse de que no exista error
	if err!=nil { // Si se produce un error ejecutamos el panic
		panic(err.Error())
	}
	// Ejecutamos la variable incertarRegistros con el metodo Exec
	insertarRegistros.Exec()


	//fmt.Fprintf(w,"Hola Go")
	// Accedemos al contenido de la plantilla inicio
	plantillas.ExecuteTemplate(w, "inicio", nil)
}

// Funcion para mostrar la plantilla de crear
func Crear(w http.ResponseWriter, r *http.Request){
	// Accedemos al contenido de la plantilla crear
	plantillas.ExecuteTemplate(w, "crear", nil)
}