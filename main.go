package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Struct base de los datos que pedires a nuestra api
type Base struct {
	ID     int    `json:"ID"`
	Nombre string `json:"Nombre"`
	Genero string `json:"Genero"`
}
type Peticiones []Base

var peticion = Peticiones{{
	ID:     1,
	Nombre: "Juan",
	Genero: "Masculino",
}}

// Funcion index
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a mi Api Rest")
}

// Funcion para crear registros
func C_Peticion(w http.ResponseWriter, r *http.Request) {
	// Creamos la variable que guarda los datos
	var N_peticion Base
	// Creamos una variable temporal para la validacion y lectura de datos
	reqBody, err := io.ReadAll(r.Body)
	// Evaluamos si se encuentra algun error en la peticion
	if err != nil {
		fmt.Fprintf(w, "Los datos ingresados no son correctos")
		return
	}

	// Asignacion de datos que son leidos por la variable body
	json.Unmarshal(reqBody, &N_peticion)
	// Incrementacion del ID
	N_peticion.ID = len(peticion) + 1
	// Añadimos el registro a nuestro arreglo
	peticion = append(peticion, N_peticion)
	// Indicamos los tipos de datos que se ingresan al server
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(N_peticion)
}

// Fincion para actualizar
func F_Actualizar(w http.ResponseWriter, r *http.Request) {
	//Variable para reciir las respuestas generales del servidor
	vars := mux.Vars(r)

	//Creamos una variable que me perimita obtener el ide
	requestID, err := strconv.Atoi(vars["id"])
	//creamos uan varaiable que guarde los datos qye se van actualizar
	var Datos Base

	if err != nil {
		fmt.Fprintf(w, "No se ha encontrado ningun registro para actualizar")
		return
	}

	//Lectura del body
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al leer los datos")
		return
	}

	//Asignamos los datos a los registros
	json.Unmarshal(reqBody, &Datos)

	//Hacemos la busqueda de ID
	for i, b := range peticion {
		//comprobacion de ID
		if b.ID == requestID {
			//Asignar el ID a la actualizacion
			Datos.ID = b.ID
			//Transferir los datos de la variable a los registros
			peticion[i] = Datos
			//Mensaje en pantalla
			fmt.Fprintf(w, "El registro %v ha sido cambiado con Exito", requestID)
			return
		}
	}

	// Si no se encuentra el ID
	http.Error(w, "Registro no encontrado", http.StatusNotFound)
}

// Funcion para eliminar los datos
func F_Eliminar(w http.ResponseWriter, r *http.Request) {
	eliminacion := mux.Vars(r)
	BusquedaID, err := strconv.Atoi(eliminacion["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, t := range peticion {
		if t.ID == BusquedaID {
			peticion = append(peticion[:i], peticion[i+1:]...)
			fmt.Fprintf(w, "El registro con ID %d se ha eliminado correctamente", BusquedaID)
			return
		}
	}

	http.Error(w, "Registro no encontrado", http.StatusNotFound)
}

// Funcion para mostrar datos individuales
func M_Individuales(w http.ResponseWriter, r *http.Request) {
	//Variable de lectura de datos
	mostrar := mux.Vars(r)
	//Variable que reciba el id
	BusquedaId, err := strconv.Atoi(mostrar["id"])
	//Evalua si existe el id
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	// Busqueda de registro
	for _, t := range peticion {
		//Evaluamos si el registro ha sido encontrado
		if t.ID == BusquedaId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(t)
			return
		}
	}
	// Si no se encuentra el registro
	http.Error(w, "Registro no encontrado", http.StatusNotFound)
}

// Funcion para mostrar todos los datos
func M_Todo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peticion)
}

// Rutas
func main() {
	ruta := mux.NewRouter().StrictSlash(true)

	// Ruta de bienvenida
	ruta.HandleFunc("/", Index)

	// Ruta para mostrar todos los datos
	ruta.HandleFunc("/Datos", M_Todo).Methods("GET")

	// Ruta para crear peticiones en la API
	ruta.HandleFunc("/Datos", C_Peticion).Methods("POST")

	// Ruta para eliminar datos
	ruta.HandleFunc("/Datos/{id}", F_Eliminar).Methods("DELETE")

	//Ruta para actualizar los datos
	ruta.HandleFunc("/Datos/{id}", F_Actualizar).Methods("PUT")

	//Ruta para busqueda individual
	ruta.HandleFunc("/Datos/{id}", M_Individuales).Methods("GET")

	// Inicializacion del servidor
	log.Fatal(http.ListenAndServe(":8000", ruta))
}
