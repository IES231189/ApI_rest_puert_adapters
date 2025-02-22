package aplication

import(
	"APICRUD/Productos/domain/entities"
	"APICRUD/Productos/domain/repositories"
)


type useCaseActualizarProducto struct{
 repository repositories.ProductoRepositories
}

func NuevoUseCaseActualizarProducto(repo repositories.ProductoRepositories)*useCaseActualizarProducto{
	return &useCaseActualizarProducto{repository:repo}
}

func (ps *useCaseActualizarProducto) ActualizarProducto(producto *entities.Producto) error{
	return ps.repository.ActualizarProducto(producto)
}


