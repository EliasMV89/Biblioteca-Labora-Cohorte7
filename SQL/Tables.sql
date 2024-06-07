CREATE TABLE Libro (
                       ID_Libro INT AUTO_INCREMENT PRIMARY KEY,
                       titulo VARCHAR(100) NOT NULL,
                       ID_Autor INT,
                       Año_Publicacion YEAR NOT NULL,
                       disponible TINYINT DEFAULT 1
);

CREATE TABLE Autor (
                       ID_Autor INT AUTO_INCREMENT PRIMARY KEY,
                       Nombre VARCHAR(100) NOT NULL,
                       Apellido VARCHAR(100) NOT NULL
);

CREATE TABLE Usuario (
                         ID_Usuario INT AUTO_INCREMENT PRIMARY KEY,
                         Nombre VARCHAR(100) NOT NULL,
                         Apellido VARCHAR(100) NOT NULL,
                         Email VARCHAR(100) NOT NULL,
                         Telefono VARCHAR(100) NOT NULL
);

CREATE TABLE Prestamo (
                          ID_Prestamo INT AUTO_INCREMENT PRIMARY KEY,
                          ID_Libro INT,
                          ID_Usuario INT,
                          Fecha_Prestamo DATE NOT NULL,
                          Fecha_Devolucion DATE,
                          FOREIGN KEY (ID_Libro) REFERENCES Libro(ID_Libro),
                          FOREIGN KEY (ID_Usuario) REFERENCES Usuario(ID_Usuario)
);

CREATE TABLE Administrador (
                               ID_Administrador INT AUTO_INCREMENT PRIMARY KEY,
                               Usuario VARCHAR(100),
                               Contraseña VARCHAR(100)
);


