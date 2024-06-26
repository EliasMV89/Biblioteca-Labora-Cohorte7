package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Prestamo struct {
	ID_Libro         int    `json:"id_libro"`
	ID_Usuario       int    `json:"id_usuario"`
	Fecha_Prestamo   string `json:"fecha_prestamo"`
	Fecha_Devolucion string `json:"fecha_devolucion,omitempty"`
}

// Realizar un nuevo préstamo
func PrestarLibro(db *sql.DB, prestamo Prestamo) error {
	var disponible int
	err := db.QueryRow("SELECT Disponible FROM Libro WHERE ID_Libro = ?", prestamo.ID_Libro).Scan(&disponible)
	if err != nil {
		log.Fatal(err)
	}
	if disponible == 1 {
		queryPrestamo := `INSERT INTO Prestamo (ID_Libro, ID_Usuario, Fecha_Prestamo) VALUES (?,?,?)`
		_, err = db.Exec(queryPrestamo, prestamo.ID_Libro, prestamo.ID_Usuario, prestamo.Fecha_Prestamo)
		if err != nil {
			log.Printf("Error al registrar el préstamo %v", err)
			return err
		}
		queryLibro := `UPDATE Libro SET Disponible = 0 WHERE ID_Libro = ?`
		_, err = db.Exec(queryLibro, prestamo.ID_Libro)
		if err != nil {
			log.Printf("Error al actualizar la disponibilidad del libro %v", err)
			return err
		}
	} else {
		fmt.Println("El libro no está disponible.")
		return nil
	}
	fmt.Println("Préstamo registrado con éxito.")
	return nil
}

// Agregar la fecha de devolución y actualizar el estado del libro
func DevolverLIbro(db *sql.DB, fecha_devolucion string, id_prestamo, id_libro int) error {
	queryPrestamo := `UPDATE Prestamo SET Fecha_Devolucion = ? WHERE ID_Prestamo = ?`
	_, err := db.Exec(queryPrestamo, fecha_devolucion, id_prestamo)
	if err != nil {
		log.Printf("Error al actualizar la fecha de devolución: %v", err)
		return err
	}

	queryLibro := `UPDATE Libro SET Disponible = 1 WHERE ID_Libro = ?`
	_, err = db.Exec(queryLibro, id_libro)
	if err != nil {
		log.Printf("Error al actualizar la disponibilidad del libro %v", err)
		return err
	}
	fmt.Println("Devolución realizada correctamente.")
	return nil
}

// Listar todos los préstamos actuales de un usuario
func ListarPrestamosUsuario(db *sql.DB, ID_Usuario int) error {
	query := `SELECT ID_Prestamo, ID_Libro, ID_Usuario FROM Prestamo WHERE ID_Usuario = ? AND Fecha_Devolucion IS NULL`
	rows, err := db.Query(query, ID_Usuario)
	if err != nil {
		log.Printf("Error al listar los préstamos del usuario: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var ID_Prestamo, ID_Libro, ID_Usuario int

		err := rows.Scan(&ID_Prestamo, &ID_Libro, &ID_Usuario)
		if err != nil {
			log.Printf("Error al leer fila: %v", err)
			return err
		}
		fmt.Printf("ID_Prestamo: %d, ID_Libro: %d, ID_Usuario: %d\n", ID_Prestamo, ID_Libro, ID_Usuario)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("Error al iterar las filas: %v", err)
		return err
	}
	fmt.Println("Préstamos listados correctamente.")
	return nil
}
