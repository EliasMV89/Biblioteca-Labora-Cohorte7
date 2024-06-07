package controllers

import (
	"Biblioteca/services"
	"Biblioteca/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreatePrestamo(w http.ResponseWriter, r *http.Request) {
	var prestamo services.Prestamo
	if err := json.NewDecoder(r.Body).Decode(&prestamo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	if err := services.PrestarLibro(db, prestamo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(prestamo)
}

func DevolverLibro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idPrestamo, err := strconv.Atoi(vars["id_prestamo"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	idLibro, err := strconv.Atoi(vars["id_libro"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var prestamo services.Prestamo
	if err := json.NewDecoder(r.Body).Decode(&prestamo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	if err := services.DevolverLIbro(db, prestamo.Fecha_Devolucion, idPrestamo, idLibro); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prestamo)
}

func ListarPrestamosUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idUsuario, err := strconv.Atoi(vars["id_usuario"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	if err := services.ListarPrestamosUsuario(db, idUsuario); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
