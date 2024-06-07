-- Buscar libros por titulo o autor
SELECT Libro.* FROM Libro
LEFT JOIN Autor ON Libro.ID_Autor = Autor.ID_Autor
WHERE Libro.titulo '%keyword'
OR Autor.Nombre LIKE '%keyword';
