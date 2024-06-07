CREATE TRIGGER VerificarStock
    ON StockLibro
    BEFORE INSERT
AS
BEGIN
    DECLARE @ID_Libro INT, @cantidad INT, @Cantidad_Copias INT;
SELECT @ID_Libro = ID_Libro, @cantidad = cantidad FROM INSERTED;
SELECT @copias_disponibles = copias_disponibles FROM libro WHERE libro_id = @libro_id;

IF @cantidad > @copias_disponibles
BEGIN
        RAISEERROR ('No hay suficientes copias disponibles.', 16, 1);
ROLLBACK;
END
END;
