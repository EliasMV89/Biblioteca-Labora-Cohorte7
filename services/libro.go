/*
package services

import (

	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

)

	type Libro struct {
		Titulo           string
		ID_Autor         int
		Anio_Publicacion string
		Disponible       int
	}

// AgregarLibro inserta un nuevo libro en la tabla Libro

	func AgregarLibro(db *sql.DB, libro Libro) error {
		// Prepara la consulta SQL para insertar un nuevo registro
		query := `INSERT INTO Libro (Titulo, ID_Autor, Año_Publicacion, Disponible) VALUES (?, ?, ?, ?)`
		// Ejecuta la consulta con los datos del usuario
		_, err := db.Exec(query, libro.Titulo, libro.ID_Autor, libro.Anio_Publicacion, libro.Disponible)
		if err != nil {
			log.Printf("Error al insertar libro: %v", err)
			return err
		}
		fmt.Println("Libro insertado correctamente.")
		return nil
	}

	func ListarLibrosDisponibles(db *sql.DB) error {
		query := `SELECT ID_Libro, Titulo FROM Libro WHERE Disponible = 1`
		rows, err := db.Query(query)
		if err != nil {
			log.Printf("Error al listar los libros disponibles: %v", err)
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var idLibro int
			var titulo string
			err := rows.Scan(&idLibro, &titulo)
			if err != nil {
				log.Printf("Error al leer fila: %v", err)
				return err
			}
			fmt.Printf("ID_Libro: %d, Titulo: %s\n", idLibro, titulo)
		}

		err = rows.Err()
		if err != nil {
			log.Printf("Error al iterar filas: %v", err)
			return err
		}

		fmt.Println("Libros listados correctamente.")
		return nil
	}

// Buscar libro por titulo o por autor

	func BuscarPorTituloOAutor(db *sql.DB, buscar string) error {
		query := `SELECT Libro.* FROM Libro LEFT JOIN Autor ON Libro.ID_Autor = Autor.ID_Autor WHERE Libro.Titulo LIKE ? OR Autor.Nombre LiKE ?`
		rows, err := db.Query(query, buscar, buscar)
		if err != nil {
			log.Printf("Error al buscar por titulo o autor: %v", err)
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var ID_Libro int
			var titulo string
			var ID_Autor int
			var Anio_Publicacion string
			var disponible int

			err := rows.Scan(&ID_Libro, &titulo, &ID_Autor, &Anio_Publicacion, &disponible)

			if err != nil {
				log.Printf("Error al leer fila: %v", err)
			}
			fmt.Printf("ID_Libro: %d, Titulo: %s, ID_Autor: %d Año_Publicacion: %d, Disponible: %d\n", ID_Libro, titulo, ID_Autor, Anio_Publicacion, disponible)
		}
		err = rows.Err()
		if err != nil {
			log.Printf("Error leer las filas: %v", err)
			return err
		}
		fmt.Println("Libros listados correctamente.")
		return nil
	}

// Listar libros mas populares en base a la cantidad de prestamos

	func ListarLibrosPopulares(db *sql.DB) error {
		query := `SELECT Libro.ID_Libro, Libro.titulo, COUNT(Prestamo.ID_Libro) AS Cantidad_Prestamos FROM Libro JOIN Prestamo ON Libro.ID_Libro = Prestamo.ID_Libro GROUP BY Libro.ID_Libro, Libro.titulo ORDER BY Cantidad_Prestamos DESC`
		rows, err := db.Query(query)
		if err != nil {
			log.Printf("Error al listar los libros mas populares: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var ID_Libro, Cantidad_Prestamos int
			var titulo string

			err := rows.Scan(&ID_Libro, &titulo, &Cantidad_Prestamos)
			if err != nil {
				log.Printf("Error al iterar las filas: %v", err)
				return err
			}
			fmt.Printf("ID_Libro: %d, Titulo: %s, Cantidad_Prestamos: %d\n", ID_Libro, titulo, Cantidad_Prestamos)
		}
		err = rows.Err()
		if err != nil {
			log.Printf("Error al iterar las filas: %v", err)
			return err
		}
		fmt.Println("Libros mas populares listados correctamente.")
		return nil
	}
*/
package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Libro struct {
	Titulo           string `json:"titulo"`
	ID_Autor         int    `json:"id_autor"`
	Anio_Publicacion string `json:"anio_publicacion"`
	Disponible       int    `json:"disponible"`
}

// AgregarLibro inserta un nuevo libro en la tabla Libro
func AgregarLibro(db *sql.DB, libro Libro) error {
	query := `INSERT INTO Libro (Titulo, ID_Autor, Año_Publicacion, Disponible) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, libro.Titulo, libro.ID_Autor, libro.Anio_Publicacion, libro.Disponible)
	if err != nil {
		log.Printf("Error al insertar libro: %v", err)
		return err
	}
	fmt.Println("Libro insertado correctamente.")
	return nil
}

func ListarLibrosDisponibles(db *sql.DB) error {
	query := `SELECT ID_Libro, Titulo FROM Libro WHERE Disponible = 1`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error al listar los libros disponibles: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var idLibro int
		var titulo string
		err := rows.Scan(&idLibro, &titulo)
		if err != nil {
			log.Printf("Error al leer fila: %v", err)
			return err
		}
		fmt.Printf("ID_Libro: %d, Titulo: %s\n", idLibro, titulo)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return err
	}

	fmt.Println("Libros listados correctamente.")
	return nil
}

func BuscarPorTituloOAutor(db *sql.DB, buscar string) error {
	query := `SELECT Libro.* FROM Libro LEFT JOIN Autor ON Libro.ID_Autor = Autor.ID_Autor WHERE Libro.Titulo LIKE ? OR Autor.Nombre LIKE ?`
	rows, err := db.Query(query, buscar, buscar)
	if err != nil {
		log.Printf("Error al buscar por titulo o autor: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var ID_Libro int
		var titulo string
		var ID_Autor int
		var Anio_Publicacion string
		var disponible int

		err := rows.Scan(&ID_Libro, &titulo, &ID_Autor, &Anio_Publicacion, &disponible)
		if err != nil {
			log.Printf("Error al leer fila: %v", err)
		}
		fmt.Printf("ID_Libro: %d, Titulo: %s, ID_Autor: %d Año_Publicacion: %s, Disponible: %d\n", ID_Libro, titulo, ID_Autor, Anio_Publicacion, disponible)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return err
	}
	fmt.Println("Libros listados correctamente.")
	return nil
}

func ListarLibrosPopulares(db *sql.DB) error {
	query := `SELECT Libro.ID_Libro, Libro.titulo, COUNT(Prestamo.ID_Libro) AS Cantidad_Prestamos FROM Libro JOIN Prestamo ON Libro.ID_Libro = Prestamo.ID_Libro GROUP BY Libro.ID_Libro, Libro.titulo ORDER BY Cantidad_Prestamos DESC`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error al listar los libros mas populares: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var ID_Libro, Cantidad_Prestamos int
		var titulo string

		err := rows.Scan(&ID_Libro, &titulo, &Cantidad_Prestamos)
		if err != nil {
			log.Printf("Error al iterar las filas: %v", err)
			return err
		}
		fmt.Printf("ID_Libro: %d, Titulo: %s, Cantidad_Prestamos: %d\n", ID_Libro, titulo, Cantidad_Prestamos)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return err
	}
	fmt.Println("Libros mas populares listados correctamente")
	return nil
}
