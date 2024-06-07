-- Crear un nuevo prestamo, suponiendo que el ID_Libro es 1 y el ID_usuario es 1
INSERT INTO Prestamo (ID_Libro, ID_Usuario, Fecha_Prestamo)
VALUES (1, 1, CURDATE());

-- Actualizar el estado del libro como no disponible en la tabla Libro
UPDATE Libro
SET disponible = 0
WHERE ID_Libro = 1;

-- Verifica si el libro está disponible
SELECT disponible FROM Libro WHERE ID_Libro = your_libro_id;

-- Si el libro está disponible (disponible = 1), realiza el préstamo y actualiza la disponibilidad del libro
INSERT INTO Prestamo (ID_Libro, ID_Usuario, Fecha_Prestamo)
VALUES (your_libro_id, your_usuario_id, CURDATE());

UPDATE Libro SET disponible = 0 WHERE ID_Libro = your_libro_id;
