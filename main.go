package main

import (
	productoApp "APICRUD/Productos/aplication"
	productoController "APICRUD/Productos/infraestructure/controller"
	productoDatabase "APICRUD/Productos/infraestructure/database"
	productoRoutes "APICRUD/Productos/infraestructure/routes"
	"time"

	ofertaApp "APICRUD/Ofertas/application"
	ofertaController "APICRUD/Ofertas/infraestructure/controller"
	ofertaDatabase "APICRUD/Ofertas/infraestructure/database"
	ofertaRoutes "APICRUD/Ofertas/infraestructure/routes"

	"APICRUD/core"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("funciona?")

	db, err := core.ConnectDB()
	if err != nil {
		log.Fatal("Error al Conectar a la BD", err)
		return
	}
	defer db.Close()

	// Productos
	prodRepo := productoDatabase.NewMysqlProductoRepository(db)
	prodService := productoApp.NuevoProductServices(prodRepo)
	prodCtrl := productoController.NewProductoController(prodService)

	// Ofertas
	ofertaRepo := ofertaDatabase.NewMysqlOfertasRepository(db)
	ofertaService := ofertaApp.NewOfertaService(ofertaRepo)
	ofertaCtrl := ofertaController.NewOfertaController(ofertaService)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	productoRoutes.RegisterRoutes(router, prodCtrl)
	ofertaRoutes.RegisterOfertaRoutes(router, ofertaCtrl)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Error al conectar el servidor", err)
	}

}
