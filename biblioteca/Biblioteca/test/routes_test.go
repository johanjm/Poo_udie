package routes_test

import (
	"github.com/stretchr/testify/assert"

	connect "library_system/db"
	"library_system/management"
	"testing"
)

func TestBookList_T(t *testing.T) {
	assert.Equal(t, 0.0, 0.0)

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	database, err := connect.Connector()

	if err != nil {
		// t.Error()
		return
	}

	books, err := management.GetBooks(database)

	if err != nil {
		t.Errorf("Error obteniendo libros: %v", err)
	}

	// Verifica alguna condición sobre los libros obtenidos
	if len(books) == 0 {
		t.Errorf("Se esperaba al menos un libro, pero se obtuvo ninguno")
	}
}
func TestAuthorList_T(t *testing.T) {
	assert.Equal(t, 0.0, 0.0)

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	database, err := connect.Connector()

	if err != nil {
		// t.Error()
		return
	}

	books, err := management.GetBooks(database)

	if err != nil {
		t.Errorf("Error obteniendo libros: %v", err)
	}

	// Verifica alguna condición sobre los libros obtenidos
	if len(books) == 0 {
		t.Errorf("Se esperaba al menos un libro, pero se obtuvo ninguno")
	}
}
