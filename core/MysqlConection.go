package core

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {
	// Cargar el archivo .env
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error cargando archivo .env: %v", err)
	}

	// Construir el DSN (Data Source Name)
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME")

	// Abrir la conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al conectar la base de datos: %v", err)
	}

	return db, nil
}
