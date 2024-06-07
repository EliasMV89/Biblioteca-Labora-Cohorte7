/*
package main

import (

	"Biblioteca/services"
	"Biblioteca/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"

)

	func main() {
		// Establece la conexión con la base de datos
		db, err := utils.ConectarBaseDeDatos()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Opciones del sistema
		for {
			fmt.Print("Bienvenido!")
			fmt.Print("*******************************************\n")
			fmt.Print("Elige una opcion\n")

			fmt.Print("*******************************************\n")
			fmt.Println("1. Agregar un nuevo libro")
			fmt.Println("2. Agregar un nuevo Autor")
			fmt.Println("3. Agregar un nuevo usuario")
			fmt.Println("4. Agregar un nuevo prestamo")
			fmt.Println("5. Registrar devolucion de un libro")
			fmt.Println("6. Listar todos los libros disponibles")
			fmt.Println("7. Buscar libro por titulo o autor")
			fmt.Println("8. Mostrar los prestamos actuales a un usuario especifico")
			fmt.Println("9. Mostrar los libros mas populares segun la cantidad de prestamos")
			fmt.Println("10. Salir del sistema")
			fmt.Println("*********************************")
			fmt.Print("Ingrese su opcion: ")
			var choice int
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				fmt.Println("Agregar un nuevo libro")
				fmt.Print("Ingrese el titulo del libro: ")
				var titulo string
				fmt.Scanln(&titulo)
				fmt.Print("Ingrese el ID del autor: ")
				var id_autor int
				fmt.Scanln(&id_autor)
				fmt.Print("Ingrese año de publicacion: ")
				var anio_publicacion string
				fmt.Scanln(&anio_publicacion)
				nuevoLibro := services.Libro{
					Titulo:           titulo,
					ID_Autor:         id_autor,
					Anio_Publicacion: anio_publicacion,
					Disponible:       1,
				}
				services.AgregarLibro(db, nuevoLibro)
			case 2:
				fmt.Println("Agregar un nuevo autor")
				fmt.Print("Ingrese el nombre del autor: ")
				var nombre string
				fmt.Scanln(&nombre)
				nuevoAutor := services.Autor{
					Nombre: nombre,
				}
				services.AgregarAutor(db, nuevoAutor)

			case 3:
				fmt.Println("Agregar un nuevo usuario")
				fmt.Print("Ingrese el nombre del usuario: ")
				var nombre string
				fmt.Scanln(&nombre)
				fmt.Print("Ingrese el apellido del usuario: ")
				var apellido string
				fmt.Scanln(&apellido)
				fmt.Print("Ingrese el email del usuario:")
				var email string
				fmt.Scanln(&email)
				fmt.Print("Ingrese el telefono del usuario: ")
				var telefono string
				fmt.Scanln(&telefono)
				nuevoUsuario := services.Usuario{
					Nombre:   nombre,
					Apellido: apellido,
					Email:    email,
					Telefono: telefono,
				}
				services.AgregarUsuario(db, nuevoUsuario)
			case 4:
				fmt.Println("Agregar un nuevo prestamo")
				fmt.Print("Ingrese el ID del libro: ")
				var id_libro int
				fmt.Scanln(&id_libro)
				fmt.Print("Ingrese el ID del usuario: ")
				var id_usuario int
				fmt.Scanln(&id_usuario)
				t := time.Now()
				var fecha_Prestamo string = t.Format("2006-01-02")
				nuevoPrestamo := services.Prestamo{
					ID_Libro:       id_libro,
					ID_Usuario:     id_usuario,
					Fecha_Prestamo: fecha_Prestamo,
				}
				services.PrestarLibro(db, nuevoPrestamo)

			case 5:
				fmt.Println("Registrar devolucion de un libro")
				fmt.Print("Ingrese el ID del prestamo: ")
				var id_prestamo int
				fmt.Scanln(&id_prestamo)
				fmt.Print("Ingrese el ID del libro: ")
				var id_libro int
				fmt.Scanln(&id_libro)
				t := time.Now()
				var fecha_devolucion string = t.Format("2006-01-02")
				services.DevolverLIbro(db, fecha_devolucion, id_prestamo, id_libro)
			case 6:
				fmt.Println("Listar todos los libros disponibles")
				services.ListarLibrosDisponibles(db)
			case 7:
				fmt.Println("Buscar libro por titulo o autor")
				fmt.Print("Ingrese el titulo o el autor: ")
				var buscar string
				fmt.Scanln(&buscar)
				services.BuscarPorTituloOAutor(db, buscar)
			case 8:
				fmt.Println("Listar prestamos actuales por usuario")
				fmt.Print("Ingrese el ID del usuario: ")
				var id_usuario int
				fmt.Scanln(&id_usuario)
				services.ListarPrestamosUsuario(db, id_usuario)
			case 9:
				fmt.Println("Listar los libros mas populares por prestamo")
				services.ListarLibrosPopulares(db)
			case 10:
				os.Exit(0)
			default:
				fmt.Println("Opcion invalida")

			}
		}

}
*/
package main

import (
	"Biblioteca/routers"
	"Biblioteca/utils"
	"log"
	"net/http"
)

func main() {
	utils.InitDB()
	router := routers.InitRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
