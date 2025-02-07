package entities

type Producto struct {
	Id_producto    int
	Nombre         string
	Descripcion    string
	Precio         float32
	Stock          int
	Imagen_url     string
	Id_categoria   int
	Fecha_creacion string
}
