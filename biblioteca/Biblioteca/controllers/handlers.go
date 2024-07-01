package controllers

import (
	"fmt"
	"html/template"
	"io"
	connect "library_system/db"
	"library_system/management"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

type Error struct {
	Message string
}

// var authors = []management.Author{
// 	{
// 		Person: management.Person{
// 			FullName: "Gabriel García Márquez",
// 		},
// 		Books: []management.Book{
// 			{
// 				Title:           "One Hundred Years of Solitude",
// 				PublicationDate: time.Date(1967, time.May, 30, 0, 0, 0, 0, time.UTC),
// 				File:            "/files/one_hundred_years_of_solitude.pdf",
// 			},
// 			{
// 				Title:           "Love in the Time of Cholera",
// 				PublicationDate: time.Date(1985, time.November, 8, 0, 0, 0, 0, time.UTC),
// 				File:            "/files/love_in_the_time_of_cholera.pdf",
// 			},
// 		},
// 	},
// 	{
// 		Person: management.Person{
// 			FullName: "J.K. Rowling",
// 		},
// 		Books: []management.Book{
// 			{
// 				Title:           "Harry Potter and the Philosopher's Stone",
// 				PublicationDate: time.Date(1997, time.June, 26, 0, 0, 0, 0, time.UTC),
// 				File:            "/files/harry_potter_and_the_philosophers_stone.pdf",
// 			},
// 			{
// 				Title:           "Harry Potter and the Chamber of Secrets",
// 				PublicationDate: time.Date(1998, time.July, 2, 0, 0, 0, 0, time.UTC),
// 				File:            "/files/harry_potter_and_the_chamber_of_secrets.pdf",
// 			},
// 		},
// 	},
// }

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	t, err := template.ParseFiles("templates/base.html", "templates/"+tmpl)
	if err != nil {
		return err
	}

	// Render template
	t.ExecuteTemplate(w, "base", data)

	return nil
}

func GetBooksByAuthor(db *gorm.DB, authorID uint) ([]management.Book, error) {
	var books []management.Book

	// Load all books by the author with the specified ID
	result := db.Where("author_id = ?", authorID).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

// Handlers

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	database, erro := connect.Connector()
	if erro != nil {
		http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
			return
		}

		// Obtener los datos del formulario
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		println("form data", email, password)

		// Validar campos (ejemplo: verificar que los campos no estén vacíos)
		if email == "" || password == "" {
			err := Error{Message: "Los campos no pueden ser vacíos!"}
			RenderTemplate(w, "login.html", err)
			return
		}

		// Consultar la base de datos para verificar las credenciales
		var user management.Account
		result := database.Where("email = ?", email).First(&user)
		if result.Error != nil {
			err := Error{Message: "El usuario no existe!"}
			RenderTemplate(w, "login.html", err)
			return
		}

		// Verificar la contraseña
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			// http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
			err := Error{Message: "La contraseña no coincide!"}
			RenderTemplate(w, "login.html", err)
			return
		}

		// http.Redirect(w, r, "/login", http.StatusSeeOther)
		RenderTemplate(w, "login.html", nil)
	} else {
		// Método GET: Mostrar el formulario de login
		RenderTemplate(w, "login.html", nil)
	}
}

// Función para manejar la solicitud POST de /signup

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	// Establecer conexión con la base de datos utilizando la función Connector de db package
	database, err := connect.Connector()

	if err != nil {
		http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "GET":
		// Mostrar el formulario de registro
		RenderTemplate(w, "signup.html", nil)

	case "POST":
		// Procesar los datos del formulario POST
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
			return
		}

		// Obtener los datos del formulario
		fullname := r.Form.Get("fullname")
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		// Validar campos (ejemplo: verificar que los campos no estén vacíos)
		if fullname == "" || email == "" || password == "" {
			http.Error(w, "Todos los campos son requeridos", http.StatusBadRequest)
			return
		}

		// Hash de la contraseña usando bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error al crear la contraseña", http.StatusInternalServerError)
			return
		}

		// Crear una nueva instancia de Account con los datos del formulario
		newAccount := management.Account{
			Email:    email,
			Password: string(hashedPassword),
			Person: management.Person{
				FullName: fullname,
			},
		}

		// Guardar en la base de datos utilizando la función CreateAccount de management package
		err = management.CreateAccount(database, &newAccount)
		if err != nil {
			http.Error(w, "Error al crear el usuario", http.StatusInternalServerError)
			return
		}

		// Redireccionar o mostrar algún mensaje de éxito
		// http.Redirect(w, r, "/login", http.StatusSeeOther)
		RenderTemplate(w, "login.html", nil)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index.html", nil)
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	database, err := connect.Connector()

	if err != nil {
		http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}

	defer func() {
		sqlDB, err := database.DB()
		if err != nil {
			fmt.Println("Error al obtener objeto sql.DB:", err)
		} else {
			sqlDB.Close()
		}
	}()

	switch r.Method {
	case http.MethodGet:
		// Consultar la lista de autores desde la base de datos
		books, err := management.GetBooks(database)
		if err != nil {
			// http.Error(w, "Error al obtener la lista de autores", http.StatusInternalServerError)
			err := Error{Message: "Error al obtener la lista de libros"}
			RenderTemplate(w, "books.html", err)
			return
		}

		// fmt.Print(books[0].Title)

		RenderTemplate(w, "books.html", books)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func AuthorsHandler(w http.ResponseWriter, r *http.Request) {
	database, err := connect.Connector()

	if err != nil {
		http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}
	// Asegurarse de cerrar la conexión de la base de datos
	defer func() {
		sqlDB, err := database.DB()
		if err != nil {
			fmt.Println("Error al obtener objeto sql.DB:", err)
		} else {
			sqlDB.Close()
		}
	}()

	switch r.Method {
	case http.MethodGet:
		// Consultar la lista de autores desde la base de datos
		authors, err := management.GetAuthors(database)
		if err != nil {
			// http.Error(w, "Error al obtener la lista de autores", http.StatusInternalServerError)
			err := Error{Message: "Error al obtener la lista de autores"}
			RenderTemplate(w, "authors.html", err)
			return
		}

		RenderTemplate(w, "authors.html", authors)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func NewAuthorHandler(w http.ResponseWriter, r *http.Request) {
	database, err := connect.Connector()

	if err != nil {
		http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		RenderTemplate(w, "new_author.html", nil)

	case http.MethodPost:
		fullname := r.FormValue("fullname")
		biography := r.FormValue("bio")
		phone := r.FormValue("phone")
		birthYearStr := r.FormValue("birth-year")
		nationality := r.FormValue("nationality")
		fmt.Println("year", birthYearStr)

		// Validar y convertir el año de nacimiento a entero
		birthDate, err := time.Parse("2006-01-02", birthYearStr)
		if err != nil {
			// msg := Error{Message: fmt.Printf("Error producido por mala fecha: %v\n", err)}
			err := Error{Message: "Error producido por mala fecha"}
			RenderTemplate(w, "new_author.html", err)
			return
		}

		// Crear un nuevo objeto Author y Person
		newPerson := management.Person{
			FullName: fullname,
			// Birthdate:   time.Date(birthDate, time.January, 1, 0, 0, 0, 0, time.UTC),
			Birthdate:   birthDate,
			Nationality: nationality,
			Phone:       phone,
		}

		newAuthor := management.Author{
			Biography: biography,
			Person:    newPerson,
			Books:     nil,
		}

		// Guardar el nuevo autor en la base de datos
		err = management.CreateAuthor(database, &newAuthor) // Pasa un puntero a newAuthor
		if err != nil {
			// http.Error(w, "Error al crear el autor", http.StatusInternalServerError)
			err := Error{Message: "Error al obtener la lista de autores"}
			RenderTemplate(w, "new_author.html", err)
			return
		}

		authors, err := management.GetAuthors(database)
		if err != nil {
			// http.Error(w, "Error al obtener la lista de autores", http.StatusInternalServerError)
			err := Error{Message: "Error al obtener la lista de autores"}
			RenderTemplate(w, "new_author.html", err)
			return
		}
		// Redirigir a una página de éxito o renderizar una plantilla que confirme la creación
		RenderTemplate(w, "authors.html", authors)
		// Redirigir a una página de éxito o renderizar una plantilla que confirme la creación
		// http.Redirect(w, r, "/authors", http.StatusSeeOther)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func DeleteAuthorHandler(w http.ResponseWriter, r *http.Request) {
	database, err := connect.Connector()

	if err != nil {
		http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	fmt.Println("id for delete", id)

	err = management.DeleteAuthor(database, id)
	if err != nil {
		http.Error(w, "Unable to delete user", http.StatusInternalServerError)
		return
	}

	authors, err := management.GetAuthors(database)
	if err != nil {
		err := Error{Message: "Error al obtener la lista de autores"}
		RenderTemplate(w, "new_author.html", err)
		return
	}
	RenderTemplate(w, "authors.html", authors)
}

func EditAuthorHandler(w http.ResponseWriter, r *http.Request) {
	database, err := connect.Connector()

	if err != nil {
		http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case "GET":
		// Obtener el autor por ID desde la base de datos
		author, err := management.GetAuthorById(database, uint(id))
		if err != nil {
			// Manejar el error al obtener el autor
			err := Error{Message: "Error al obtener el autor"}
			RenderTemplate(w, "edit_author.html", err)
			return
		}
		RenderTemplate(w, "edit_author.html", author)

	case "POST":
		// Obtener los datos del formulario
		fullname := r.FormValue("fullname")
		biography := r.FormValue("bio")
		phone := r.FormValue("phone")
		birthYearStr := r.FormValue("birth-year")
		nationality := r.FormValue("nationality")

		// Validar y convertir el año de nacimiento a entero
		birthDate, err := time.Parse("2006-01-02", birthYearStr)
		if err != nil {
			http.Error(w, "Error producido por mala fecha", http.StatusBadRequest)
			return
			// // msg := Error{Message: fmt.Printf("Error producido por mala fecha: %v\n", err)}
			// err := Error{Message: "Error producido por mala fecha"}
			// RenderTemplate(w, "edit_author.html", err)
			// return
		}

		// Crear un nuevo objeto Author con Person integrado
		newAuthor := management.Author{
			Biography: biography,
			Person: management.Person{
				FullName:    fullname,
				Birthdate:   birthDate,
				Nationality: nationality,
				Phone:       phone,
			},
		}

		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		err = management.UpdateAuthor(database, ID, &newAuthor)
		if err != nil {
			http.Error(w, "Error al crear el autor", http.StatusInternalServerError)
			// return
			// err := Error{Message: "Error al actualizar el autor"}
			// RenderTemplate(w, "edit_author.html", err)
		}

		authors, err := management.GetAuthors(database)
		if err != nil {
			http.Error(w, "Error al obtener la lista de autores", http.StatusInternalServerError)
			// err := Error{Message: "Error al obtener la lista de autores"}
			// RenderTemplate(w, "edit_author.html", err)
			return
		}
		// Redirigir a una página de éxito o renderizar una plantilla que confirme la creación
		RenderTemplate(w, "authors.html", authors)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func EditPersonHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "edit_person.html", nil)
}

func SaveBookHandler(w http.ResponseWriter, r *http.Request) {
	database, err := connect.Connector()

	if err != nil {
		http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		authors, err := management.GetAuthors(database)
		if err != nil {
			// http.Error(w, "Error al obtener la lista de autores", http.StatusInternalServerError)
			err := Error{Message: "Error al obtener la lista de autores"}
			RenderTemplate(w, "authors.html", err)
			return
		}

		RenderTemplate(w, "new_book.html", authors)
	case http.MethodPost:
		title := r.FormValue("title")
		publishedDate := r.FormValue("publishedDate")
		barcode := r.FormValue("barcode")
		// pageCountStr := r.FormValue("pageCount")
		// format := r.FormValue("format")
		// dueDate := r.FormValue("dueDate")
		authorIDStr := r.FormValue("authorID")

		// Leer el archivo subido
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error al leer el archivo", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Leer el contenido del archivo en un []byte
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Error al leer el archivo", http.StatusInternalServerError)
			return
		}

		// Convertir el string a int
		authorIDInt, err := strconv.Atoi(authorIDStr)
		if err != nil {
			http.Error(w, "ID del autor inválido", http.StatusBadRequest)
			return
		}

		// Convertir int a uint
		authorID := uint(authorIDInt)

		// Convertir la fecha de publicación a time.Time
		dateBook, err := time.Parse("2006-01-02", publishedDate)
		if err != nil {
			err := Error{Message: "Error producido por mala fecha del libro"}
			RenderTemplate(w, "new_book.html", err)
			return
		}

		// Convertir pageCount a int
		// pageCount, err := strconv.Atoi(pageCountStr)
		// if err != nil {
		// 	http.Error(w, "Conteo de páginas inválido", http.StatusBadRequest)
		// 	return
		// }

		// Convertir la fecha de vencimiento a time.Time
		// dueDateBook, err := time.Parse("2006-01-02", dueDate)
		// if err != nil {
		// 	err := Error{Message: "Error producido por mala fecha de vencimiento"}
		// 	RenderTemplate(w, "new_book.html", err)
		// 	return
		// }

		// Crear un nuevo objeto Book
		book := management.Book{
			Title:           title,
			PublicationDate: dateBook,
			File:            fileBytes,
			AuthorID:        authorID,
			Barcode:         barcode,
			// PageCount:       pageCount,
			// Format:          format,
			// DueDate:         dueDateBook,
		}

		// Guardar el nuevo libro en la base de datos
		err = management.CreateBook(database, &book)
		if err != nil {
			err := Error{Message: "Error al crear el libro"}
			RenderTemplate(w, "new_book.html", err)
			return
		}

		authors, err := management.GetAuthors(database)
		if err != nil {
			err := Error{Message: "Error al obtener la lista de autores"}
			RenderTemplate(w, "new_author.html", err)
			return
		}

		// Redirigir a una página de éxito o renderizar una plantilla que confirme la creación
		RenderTemplate(w, "authors.html", authors)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}

}

func EditBookHandler(w http.ResponseWriter, r *http.Request) {
	database, err := connect.Connector()

	if err != nil {
		http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}

	defer func() {
		sqlDB, err := database.DB()
		if err != nil {
			fmt.Println("Error al obtener objeto sql.DB:", err)
		} else {
			sqlDB.Close()
		}
	}()

}
