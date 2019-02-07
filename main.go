package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Amigo struct {
	Nombre	string `json:"nombre"`
	Email	string `json:"email"`
	Edad	int `json:"edad"`
	Telefono int `json:"telefono"`
}

/*
Se usa postman
Crear amigos
listar amigos ( por nombre, alfabetico )
borrar un amigo
editar un amigo
*/
type Amigos []Amigo

// Los handlers son los manejadores para las rutas de navegacion
// Los hadlersfunc son un tipo de dato que permite convertir funciones en handlers
func handleRequests() {
	// route
	http.HandleFunc("/listar", listar)
	http.HandleFunc("/crear", crear)
	http.HandleFunc("/editar", editar)
	http.HandleFunc("/borrar", borrar)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func homePage(w http.ResponseWriter, r *http.Request) { //Donde se devuelve la respuesta al cliente / la peticion del cliente
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {

	handleRequests()
}

func listar(w http.ResponseWriter, r *http.Request) {
	amigos := Amigos{
		Amigo{Nombre: "Diego", Email: "chavezddiego@gmail.com", Edad: 21, Telefono: 45678987},
		Amigo{Nombre: "Marcos", Email: "marcos@gmail.com", Edad: 30, Telefono: 97818987},
		Amigo{Nombre: "Juan", Email: "marcos@gmail.com", Edad: 30, Telefono: 97818987},
	}
	fmt.Println(amigos)

	json.NewEncoder(w).Encode(amigos)
}

func crear(w http.ResponseWriter, r *http.Request) {
}

func editar(w http.ResponseWriter, r *http.Request) {
}

func borrar(w http.ResponseWriter, r *http.Request) {
}