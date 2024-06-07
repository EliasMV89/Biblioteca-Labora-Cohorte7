package controllers

import (
	"Biblioteca/services"
	"Biblioteca/utils"
	"encoding/json"
	"net/http"
)

func GetLibros(w http.ResponseWriter, r *http.Request) {
	// Implementación pendiente para obtener todos los libros
}

func GetLibro(w http.ResponseWriter, r *http.Request) {
	// Implementación pendiente para obtener un libro por ID
}

func CreateLibro(w http.ResponseWriter, r *http.Request) {
	var libro services.Libro
	if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	if err := services.AgregarLibro(db, libro); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(libro)
}

func ListarLibrosDisponibles(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	if err := services.ListarLibrosDisponibles(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func BuscarPorTituloOAutor(w http.ResponseWriter, r *http.Request) {
	buscar := r.URL.Query().Get("buscar")
	db := utils.GetDB()
	if err := services.BuscarPorTituloOAutor(db, buscar); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ListarLibrosPopulares(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	if err := services.ListarLibrosPopulares(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Agrega más manejadores (Update, Delete) según sea necesario
