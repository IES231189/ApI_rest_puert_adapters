package aplication

import (
	"APICRUD/Productos/domain/entities"
	"APICRUD/Productos/domain/repositories"
)

type ProductoServices struct {
	repository repositories.ProductoRepositories
}

func NuevoProductServices(repo repositories.ProductoRepositories) *ProductoServices {
	return &ProductoServices{repository: repo}
}

func (ps *ProductoServices) ListarProductos() ([]*entities.Producto, error) {
	return ps.repository.MostrarProductos()
}

func (ps *ProductoServices) AñadirProducto(producto *entities.Producto) error {
	return ps.repository.AgregarProducto(producto)
}

func (ps *ProductoServices) ActualizarProducto(producto *entities.Producto) error {
	return ps.repository.ActualizarProducto(producto)
}

func (ps *ProductoServices) EliminarProducto(id int) error {
	return ps.repository.EliminarProducto(id)
}
