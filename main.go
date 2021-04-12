package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Ticket struct {
	Id                 int
	Usuario            string
	Descripcion        string
	FechaCreacion      string
	FechaActualizacion string
	NumTicket          int
}

func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := "admin"
	Nombre := "tickets"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1:3306)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/buscar", Buscar)
	http.HandleFunc("/consulta", Consulta)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insertar", Insertar)
	http.HandleFunc("/borrar", Borrar)
	http.HandleFunc("/editar", Editar)
	http.HandleFunc("/show", Show)

	http.HandleFunc("/actualizar", Actualizar)
	log.Println("Servidor RUN")
	http.ListenAndServe(":8080", nil)
}

func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}

func Insertar(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		usuario := r.FormValue("Usuario")
		fecha := r.FormValue("Fecha")
		fechaAct := r.FormValue("Fecha")
		descripcion := r.FormValue("Descripcion")
		numTicket := ConsultarID()
		conexionEstablecida := conexionBD()
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO tickets (usuario,descripcion,fechaCreacion,fechaActualizacion,numTicket) VALUES (?,?,?,?,?)")

		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(usuario, descripcion, fecha, fechaAct, numTicket)
		http.Redirect(w, r, "/", 301)

	}

}

func Actualizar(w http.ResponseWriter, r *http.Request) {
	fmt.Print("actualizar")

	if r.Method == "POST" {

		id := r.FormValue("Id")
		usuario := r.FormValue("Usuario")
		fecha := r.FormValue("Fecha")
		descripcion := r.FormValue("Descripcion")

		conexionEstablecida := conexionBD()
		insertarRegistros, err := conexionEstablecida.Prepare("UPDATE tickets SET usuario=?, fechaActualizacion=?, descripcion=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(usuario, fecha, descripcion, id)
		http.Redirect(w, r, "/", 301)

	}

}

func ConsultarID() int {

	conexionEstablecida := conexionBD()
	lastID, err := conexionEstablecida.Query("SELECT MAX(id) FROM tickets")

	if err != nil {
		panic(err.Error())
	}
	ticket := Ticket{}
	for lastID.Next() {
		var id int

		err = lastID.Scan(&id)
		if err != nil {
			panic(err.Error())
		}
		ticket.Id = id

	}

	return ticket.Id

}

func Consulta(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		usuario := r.FormValue("Usuario")
		conexionEstablecida := conexionBD()

		registro, err := conexionEstablecida.Query("SELECT * FROM tickets WHERE id=?", usuario)
		ticket := Ticket{}
		for registro.Next() {
			var id int
			var usuario string
			var descripcion string
			var fechaCreacion string
			var fechaActualizacion string
			var numTicket int

			err = registro.Scan(&id, &usuario, &descripcion, &fechaCreacion, &fechaActualizacion, &numTicket)
			if err != nil {
				panic(err.Error())
			}
			ticket.Id = id
			ticket.Usuario = usuario
			ticket.Descripcion = descripcion
			ticket.FechaCreacion = fechaCreacion
			ticket.FechaActualizacion = fechaActualizacion
			ticket.NumTicket = numTicket
		}
		fmt.Print(ticket)
		plantillas.ExecuteTemplate(w, "show", ticket)
	}

}

func Show(w http.ResponseWriter, r *http.Request) {

	idTicker := r.URL.Query().Get("id")
	conexionEstablecida := conexionBD()
	registro, err := conexionEstablecida.Query("SELECT * FROM tickets WHERE id=?", idTicker)

	if err != nil {
		panic(err.Error())
	}
	ticket := Ticket{}
	for registro.Next() {
		var id int
		var usuario string
		var descripcion string
		var fechaCreacion string
		var fechaActualizacion string
		var numTicket int

		err = registro.Scan(&id, &usuario, &descripcion, &fechaCreacion, &fechaActualizacion, &numTicket)
		if err != nil {
			panic(err.Error())
		}
		ticket.Id = id
		ticket.Usuario = usuario
		ticket.Descripcion = descripcion
		ticket.FechaCreacion = fechaCreacion
		ticket.FechaActualizacion = fechaActualizacion
		ticket.NumTicket = numTicket
	}

	plantillas.ExecuteTemplate(w, "show", ticket)

}

func Borrar(w http.ResponseWriter, r *http.Request) {
	idTicker := r.URL.Query().Get("id")
	conexionEstablecida := conexionBD()
	borrarRegistros, err := conexionEstablecida.Prepare("DELETE FROM tickets WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	borrarRegistros.Exec(idTicker)
	http.Redirect(w, r, "/", 301)

}

func Editar(w http.ResponseWriter, r *http.Request) {
	idTicker := r.URL.Query().Get("id")
	conexionEstablecida := conexionBD()
	registro, err := conexionEstablecida.Query("SELECT * FROM tickets WHERE id=?", idTicker)

	if err != nil {
		panic(err.Error())
	}
	ticket := Ticket{}
	for registro.Next() {
		var id int
		var usuario string
		var descripcion string
		var fechaCreacion string
		var fechaActualizacion string
		var numTicket int

		err = registro.Scan(&id, &usuario, &descripcion, &fechaCreacion, &fechaActualizacion, &numTicket)
		if err != nil {
			panic(err.Error())
		}
		ticket.Id = id
		ticket.Usuario = usuario
		ticket.Descripcion = descripcion
		ticket.FechaCreacion = fechaCreacion
		ticket.FechaActualizacion = fechaActualizacion
		ticket.NumTicket = numTicket
	}

	plantillas.ExecuteTemplate(w, "editar", ticket)
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := conexionBD()
	registros, err := conexionEstablecida.Query("SELECT * FROM tickets")

	if err != nil {
		panic(err.Error())
	}

	ticket := Ticket{}
	arregloTicket := []Ticket{}

	for registros.Next() {
		var id int
		var usuario string
		var descripcion string
		var fechaCreacion string
		var fechaActualizacion string
		var numTicket int

		err = registros.Scan(&id, &usuario, &descripcion, &fechaCreacion, &fechaActualizacion, &numTicket)
		if err != nil {
			panic(err.Error())
		}
		ticket.Id = id
		ticket.Usuario = usuario
		ticket.Descripcion = descripcion
		ticket.FechaCreacion = fechaCreacion
		ticket.FechaActualizacion = fechaActualizacion
		ticket.NumTicket = numTicket

		arregloTicket = append(arregloTicket, ticket)
	}

	plantillas.ExecuteTemplate(w, "inicio", arregloTicket)

}

func Buscar(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "buscar", nil)

}
