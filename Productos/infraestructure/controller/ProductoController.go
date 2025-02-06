package controller

import (
	"APICRUD/Productos/aplication"
	"APICRUD/Productos/domain/entities"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
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

func (pc *ProductoController) SubirImagen(c *gin.Context) {
	// Obtener el archivo desde la solicitud
	file, _ := c.FormFile("imagen")
	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se proporcionó ninguna imagen"})
		return
	}

	extension := filepath.Ext(file.Filename)
	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d%s", timestamp, extension)

	// Guardar el archivo en la carpeta uploads
	savePath := filepath.Join("uploads", fileName)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la imagen"})
		return
	}

	// Retornar la URL de la imagen para almacenarla en la base de datos
	c.JSON(http.StatusOK, gin.H{"imagen_url": savePath})
}
