package routers

import (
	"Biblioteca/controllers" // Ajusta el import según tu estructura
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	// Rutas para autores
	router.HandleFunc("/autores", controllers.CreateAutor).Methods("POST")
	// Rutas para libros
	router.HandleFunc("/libros", controllers.CreateLibro).Methods("POST")
	router.HandleFunc("/libros/disponibles", controllers.ListarLibrosDisponibles).Methods("GET")
	router.HandleFunc("/libros/buscar", controllers.BuscarPorTituloOAutor).Methods("GET")
	router.HandleFunc("/libros/populares", controllers.ListarLibrosPopulares).Methods("GET")
	// Rutas para préstamos
	router.HandleFunc("/prestamos", controllers.CreatePrestamo).Methods("POST")
	router.HandleFunc("/prestamos/{id_prestamo}/{id_libro}/devolver", controllers.DevolverLibro).Methods("PUT")
	router.HandleFunc("/prestamos/usuario/{id_usuario}", controllers.ListarPrestamosUsuario).Methods("GET")
	// Rutas para usuarios
	router.HandleFunc("/usuarios", controllers.CreateUsuario).Methods("POST")
	// Agrega más rutas (GET, PUT, DELETE) según sea necesario
	return router
}
