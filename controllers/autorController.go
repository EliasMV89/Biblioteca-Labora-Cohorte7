package controllers

import (
	"Biblioteca/services"
	"Biblioteca/utils"
	"encoding/json"
	"net/http"
)

func GetAutores(w http.ResponseWriter, r *http.Request) {
	// Implementación pendiente para obtener todos los autores
}

func GetAutor(w http.ResponseWriter, r *http.Request) {
	// Implementación pendiente para obtener un autor por ID
}

func CreateAutor(w http.ResponseWriter, r *http.Request) {
	var autor services.Autor
	if err := json.NewDecoder(r.Body).Decode(&autor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	if err := services.AgregarAutor(db, autor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(autor)
}

// Agrega más manejadores (Update, Delete) según sea necesario
