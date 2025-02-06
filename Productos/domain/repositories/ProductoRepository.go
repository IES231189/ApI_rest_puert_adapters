package repositories

import(
	"APICRUD/Productos/domain/entities"
)


type ProductoRepositories interface{
	MostrarProductos() ([]*entities.Producto , error)
	AgregarProducto(producto *entities.Producto) error
	ActualizarProducto(producto *entities.Producto) error
	EliminarProducto(id int)error
	BuscarPorID(id int)([]*entities.Producto , error) 
}