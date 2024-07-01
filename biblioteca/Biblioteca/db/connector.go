package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connector establece la conexión a la base de datos utilizando Gorm y devuelve el objeto DB
func Connector() (*gorm.DB, error) {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error cargando variables de entorno: %v", err)
	}

	// Configurar cadena de conexión
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Conectar a la base de datos utilizando Gorm
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %s", err)
	}

	// Comprobar la conexión
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error al obtener el objeto DB: %s", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("error al hacer ping a la base de datos: %s", err)
	}

	fmt.Println("Database has been connected!")

	return db, nil
}
