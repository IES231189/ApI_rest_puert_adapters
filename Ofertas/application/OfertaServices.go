package application

import(
	
	"APICRUD/Ofertas/domain/repositories"
	"APICRUD/Ofertas/domain/entities"

)


type OfertaServices struct{
	of repositories.OfertasRepository
}

func NewOfertaService(of repositories.OfertasRepository)*OfertaServices {
	return &OfertaServices{of:of}
} 

func(Os *OfertaServices) MostrarOfertas() ([]*entities.Ofertas , error){
	return Os.of.MostrarOfertas()
} 

func (Os *OfertaServices) CrearOfertas(oferta *entities.Ofertas) error{
		return Os.of.CrearOfertas(oferta)
}


func (Os *OfertaServices) Actualizar(oferta *entities.Ofertas) error{
	return Os.of.Actualizar(oferta)
}

func (Os *OfertaServices) Eliminar( id int) error{
	return Os.of.Eliminar(id)
}


func (Os *OfertaServices) MostrarPorID(id int)([]*entities.Ofertas , error){
	return Os.of.MostrarPorID(id)
}

func (Os *OfertaServices) VerificarNuevasOfertas() ([]*entities.Ofertas, error) {
	return Os.of.VerificarNuevasOfertas()
}

func (Os *OfertaServices) VerificarOfertasCaducadas() ([]*entities.Ofertas, error) {
	return Os.of.VerificarOfertasCaducadas()
}




