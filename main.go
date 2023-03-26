package main

import (
	"database/sql" // Modulo necesario para trabajar con sql
	"fmt"
	"log"      // Nos permite crear diferentes tipos de loggers, usando el método New
	"net/http" // Permite construir servidores HTTP
	//"fmt" 	   		// Formato de entrada y salida de datos
	"text/template" // Para trabajar con templates

	_ "github.com/go-sql-driver/mysql" // Driver para la coneccion con la base de datos
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
	// Solicitud para incertar los datos
	http.HandleFunc("/insertar", Insertar)



	// Log que indica por consola que el servidor esta corriendo 
	log.Println("Servidor corriendo...")
	// Indicamos el servidor en el que estara corriendo la aplicacion
	http.ListenAndServe(":8080", nil)
}

// Estructura para depositar los datos de los empleados
type Empleado struct {
	Id int
	Nombre string
	Correo string
}

// Funcion para mostrar la plantilla de Inicio
func Inicio(w http.ResponseWriter, r *http.Request){

	// Conexion con la base de datos
	conexionEstablecida:= conexionBD()
	registros,err:= conexionEstablecida.Query("SELECT * FROM empleados")

	// Para ejecutar la variable insertarRegistros, primero hay que asegurarse de que no exista error
	if err!=nil { // Si se produce un error ejecutamos el panic
		panic(err.Error())
	}

	empleado:= Empleado{}
	arregloEmpleado:=[]Empleado{}

	for registros.Next(){
		var id int
		var nombre, correo string
		err= registros.Scan(&id,&nombre,&correo)

		if err != nil {
			panic(err.Error())
		}
		empleado.Id= id
		empleado.Nombre= nombre
		empleado.Correo= correo

		arregloEmpleado=append(arregloEmpleado, empleado)
	}
	fmt.Println(arregloEmpleado)

	//fmt.Fprintf(w,"Hola Go")
	// Accedemos al contenido de la plantilla inicio
	plantillas.ExecuteTemplate(w, "inicio", nil)
}

// Funcion para mostrar la plantilla de crear
func Crear(w http.ResponseWriter, r *http.Request){
	// Accedemos al contenido de la plantilla crear
	plantillas.ExecuteTemplate(w, "crear", nil)
}

// Funcion para la recepcion de datos
func Insertar(w http.ResponseWriter, r *http.Request){
	if r.Method=="POST"{ // Si existe un metodo POST
		// Entonces vamos a recepcionar esos datos
		nombre:= r.FormValue("nombre")
		correo:= r.FormValue("correo")

	// Conexion con la base de datos
	conexionEstablecida:= conexionBD()
	insertarRegistros,err:= conexionEstablecida.Prepare("INSERT INTO empleados(nombre,correo) VALUES(?,?)")

	// Para ejecutar la variable insertarRegistros, primero hay que asegurarse de que no exista error
	if err!=nil { // Si se produce un error ejecutamos el panic
		panic(err.Error())
	}
	// Ejecutamos la variable incertarRegistros con el metodo Exec
	// Pasamos como parametros el nombre y el correo
	insertarRegistros.Exec(nombre,correo)

	http.Redirect(w,r,"/",301)
	}
}