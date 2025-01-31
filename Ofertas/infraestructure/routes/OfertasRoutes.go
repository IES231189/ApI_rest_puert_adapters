package routes

import (
	"APICRUD/Ofertas/infraestructure/controller"

	"github.com/gin-gonic/gin"
)

func RegisterOfertaRoutes(router *gin.Engine, ofertaController *controller.OfertaController) {
	ofertasGroup := router.Group("/ofertas")
	{
		ofertasGroup.GET("/", ofertaController.MostrarOfertas)
		ofertasGroup.POST("/", ofertaController.CrearOfertas)
		ofertasGroup.PUT("/", ofertaController.ActualizarOferta)
		ofertasGroup.DELETE("/:id", ofertaController.EliminarOferta)
	}
}
