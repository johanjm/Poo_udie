package db

import "library_system/management"

// Connector establece la conexión a la base de datos utilizando Gorm y devuelve el objeto DB
func Migrate() {
	database, err := Connector()
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&management.Account{})
	database.AutoMigrate(&management.Author{})
	database.AutoMigrate(&management.Book{})
	database.AutoMigrate(&management.BookItem{})
	database.AutoMigrate(&management.Person{})
}
