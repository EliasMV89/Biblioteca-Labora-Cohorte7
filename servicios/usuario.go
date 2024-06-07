package servicios

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Usuario struct {
	Nombre   string
	Apellido string
	Email    string
	Telefono string
}

// AgregarUsuario inserta un nuevo usuario en la base de datos
func AgregarUsuario(db *sql.DB, usuario Usuario) error {
	// Prepara la consulta SQL para insertar un nuevo registro
	query := `INSERT INTO Usuario (Nombre, Apellido, Email, Telefono) VALUES (?, ?, ?, ?)`
	// Ejecuta la consulta con los datos del usuario
	_, err := db.Exec(query, usuario.Nombre, usuario.Apellido, usuario.Email, usuario.Telefono)
	if err != nil {
		log.Printf("Error al insertar usuario: %v", err)
		return err
	}
	fmt.Println("Usuario insertado correctamente.")
	return nil
}
