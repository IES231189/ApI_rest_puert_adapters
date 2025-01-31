package controller

import (
	"APICRUD/Ofertas/application"
	"APICRUD/Ofertas/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OfertaController struct {
	service *application.OfertaServices
}

func NewOfertaController(service *application.OfertaServices) *OfertaController {
	return &OfertaController{service: service}
}

func (c *OfertaController) MostrarOfertas(ctx *gin.Context) {
	ofertas, err := c.service.MostrarOfertas()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ofertas)
}

func (c *OfertaController) CrearOfertas(ctx *gin.Context) {
	var oferta entities.Ofertas
	if err := ctx.ShouldBindJSON(&oferta); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err := c.service.CrearOfertas(&oferta)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Oferta creada correctamente"})
}

func (c *OfertaController) ActualizarOferta(ctx *gin.Context) {
	var oferta entities.Ofertas
	if err := ctx.ShouldBindJSON(&oferta); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err := c.service.Actualizar(&oferta)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Oferta actualizada correctamente"})
}

func (c *OfertaController) EliminarOferta(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = c.service.Eliminar(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Oferta eliminada correctamente"})
}
