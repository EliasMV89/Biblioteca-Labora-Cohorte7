-- Listar todos los prestamos actuales a un usuario especifico
SELECT ID_Prestamo, ID_Libro, ID_Usuario
FROM Prestamo
WHERE Fecha_Devolucion IS NULL;