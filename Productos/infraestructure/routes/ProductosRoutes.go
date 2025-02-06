package routes

import (
	"APICRUD/Productos/infraestructure/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, productoController *controller.ProductoController) {
	routes := router.Group("/productos")
	{
		routes.GET("/", productoController.ListarProductos)
		routes.POST("/", productoController.AñadirProducto)
		routes.PUT("/", productoController.ActualizarProducto)
		routes.DELETE("/:id", productoController.EliminarProducto)
		routes.GET("/:id", productoController.BuscarPorID)
		routes.POST("/subir-imagen", productoController.SubirImagen)
	}
}
