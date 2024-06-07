package controllers

import (
	"Biblioteca/services"
	"Biblioteca/utils"
	"encoding/json"
	"net/http"
)

func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario services.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	if err := services.AgregarUsuario(db, usuario); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}
