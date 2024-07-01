//@autor: Erick Andrade
//@version: 2.0
//@fecha: 23/06/2024
//@description: Library Students Application for Oriented programming language

package main

import (
	"fmt"
	"log"
	"net/http"

	"library_system/controllers"
	connect "library_system/db"

	"github.com/gorilla/mux"
)

func main() {
	// Conexión de la base de datos y migración
	connect.Migrate()

	// Configuración de las rutas y el servidor web
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.IndexHandler).Methods("GET")
	router.HandleFunc("/login", controllers.LoginHandler).Methods("GET", "POST")
	router.HandleFunc("/signup", controllers.SignupHandler).Methods("GET", "POST")

	router.HandleFunc("/authors", controllers.AuthorsHandler).Methods("GET")
	router.HandleFunc("/books", controllers.BooksHandler).Methods("GET")

	router.HandleFunc("/people/edit", controllers.EditPersonHandler).Methods("GET", "POST")

	router.HandleFunc("/author/new", controllers.NewAuthorHandler).Methods("GET", "POST")
	router.HandleFunc("/author/edit/{id}", controllers.EditAuthorHandler).Methods("GET", "POST")
	router.HandleFunc("/author/delete/{id}", controllers.DeleteAuthorHandler).Methods("GET")

	router.HandleFunc("/book/new", controllers.SaveBookHandler).Methods("GET", "POST")
	router.HandleFunc("/book/edit/{id}", controllers.EditBookHandler).Methods("GET", "POST")

	// Rutas para el control de errores y metodos no permitidos en la aplicación
	router.HandleFunc("/error", controllers.ErrorHandler).Methods("GET")
	router.HandleFunc("/not-found", controllers.NotFoundHandler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(controllers.NotFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(controllers.MethodNotAllowed)

	// Carpeta de archivos estáticos (CSS, JS, imágenes, etc.)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Configurar el servidor HTTP
	http.Handle("/", router)
	// Iniciar el servidor
	port := 8100
	log.Printf("Server listen to http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
