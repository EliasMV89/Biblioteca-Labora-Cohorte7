/*package services

import (
	"database/sql"
	"fmt"
	"log"
)

type Autor struct {
	Nombre string
}

// AgregarAutor registrar un nuevo autor en la tabla Autor
func AgregarAutor(db *sql.DB, autor Autor) error {
	// Prepara la consulta SQL para insertar un nuevo registro
	query := `INSERT INTO Autor (Nombre) VALUES (?)`
	// Ejecuta la consulta con los datos del usuario
	_, err := db.Exec(query, autor.Nombre)
	if err != nil {
		log.Printf("Error al registrar el nuevo autor: %v", err)
		return err
	}
	fmt.Println("Autor registrado correctamente.")
	return nil
}

*/

package services

import (
	"database/sql"
	"fmt"
	"log"
)

type Autor struct {
	Nombre string `json:"nombre"`
}

// AgregarAutor registra un nuevo autor en la tabla Autor
func AgregarAutor(db *sql.DB, autor Autor) error {
	query := `INSERT INTO Autor (Nombre) VALUES (?)`
	_, err := db.Exec(query, autor.Nombre)
	if err != nil {
		log.Printf("Error al registrar el nuevo autor: %v", err)
		return err
	}
	fmt.Println("Autor registrado correctamente.")
	return nil
}
