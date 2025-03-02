package aplication


import(
	"APICRUD/Productos/domain/entities"
	"APICRUD/Productos/domain/repositories"
)

type UseCaseCrearProducto struct{
	repository repositories.ProductoRepositories
}

func NuevoUseCaseCreaProducto(repo repositories.ProductoRepositories)*UseCaseCrearProducto{
	return &UseCaseCrearProducto{repository: repo}
} 

func (ps *UseCaseCrearProducto) AñadirProducto(producto *entities.Producto) error{
	return ps.repository.AgregarProducto(producto)
}