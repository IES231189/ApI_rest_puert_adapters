package controller

import (
	"APICRUD/Productos/aplication"
	"APICRUD/Productos/domain/entities"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductoController struct {
	service *aplication.ProductoServices
}

func NewProductoController(service *aplication.ProductoServices) *ProductoController {
	return &ProductoController{service: service}
}

func (pc *ProductoController) ListarProductos(c *gin.Context) {
	productos, err := pc.service.ListarProductos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productos)

}

func (pc *ProductoController) AñadirProducto(c *gin.Context) {
	var producto entities.Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "entrada no valida"})
		return
	}

	err := pc.service.AñadirProducto(&producto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Producto añadido"})

}

func (pc *ProductoController) ActualizarProducto(c *gin.Context) {
	var producto entities.Producto

	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Entrada no valida"})
		return
	}

	err := pc.service.ActualizarProducto(&producto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Actualizado"})
}

func (pc *ProductoController) EliminarProducto(c *gin.Context) {
	id := c.Param("id")

	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("no se pudo convertir el valor a entero")
		return
	}

	err = pc.service.EliminarProducto(num)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Eliminado"})
}

func (pc *ProductoController) BuscarPorID(c *gin.Context) {

	id := c.Param("id")

	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
		return
	}

	product, err := pc.service.BuscarPorID(num)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"producto": product})
}
