package repositories

import (
	"APICRUD/Ofertas/domain/entities"

)


type OfertasRepository interface{
	MostrarOfertas()([]*entities.Ofertas , error)
	CrearOfertas(of * entities.Ofertas)error
	Actualizar(of * entities.Ofertas) error
	Eliminar(id int ) error
	MostrarPorID(id int)([]*entities.Ofertas , error)
	VerificarNuevasOfertas() ([]*entities.Ofertas, error)
    VerificarOfertasCaducadas() ([]*entities.Ofertas, error)

}