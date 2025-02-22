package aplication

import(
	"APICRUD/Productos/domain/entities"
	"APICRUD/Productos/domain/repositories"
)

type UseCaseEliminarProducto struct{
	repository repositories.ProductoRepositories
}

func NuevoUseCaseEliminarProducto(repo repositories.ProductoRepositories)*UseCaseEliminarProducto{
	return &UseCaseEliminarProducto{repository:repo}
}


func(ps *UseCaseEliminarProducto) EliminarProducto(id int) error{
	return ps.repository.EliminarProducto(id)
}
