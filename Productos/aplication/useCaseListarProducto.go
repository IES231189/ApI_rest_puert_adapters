package aplication

import(
	"APICRUD/Productos/domain/entities"
	"APICRUD/Productos/domain/repositories"
)


type useCaseListarProducto struct{
	repository repositories.ProductoRepositories
}

func NuevoProductoUseCase(repo repositories.ProductoRepositories) *useCaseListarProducto{
	return &useCaseListarProducto{repository : repo}
}

func (ps *useCaseListarProducto) ListarProductos() ([]*entities.Producto , error){
	return ps.repository.MostrarProductos()
}


