-- Listar libros mas populares en base a la cantidad de prestamos

SELECT Libro.ID_Libro, Libro.titulo, COUNT(Prestamo.ID_Libro) AS Cantidad_Prestamos
FROM Libro
JOIN Prestamo ON Libro.ID_Libro = Prestamo.ID_Libro
GROUP BY Libro.ID_Libro, Libro.titulo
ORDER BY Cantidad_Prestamos DESC;