# Gestión de Biblioteca - Endpoints

### INTEGRANTES
- MARCELO
- ERICK ANDRADE
- JOHAN QUINATOA

A continuación se detallan los endpoints disponibles para la aplicación de gestión de biblioteca desarrollada en Go:

## Endpoints

- **Página Principal**
  - **URL:** `http://localhost:8100/`
  - **Descripción:** Página principal que muestra enlaces a las secciones de autores y libros.
  - **Método:** `GET`
  - **Handler:** `IndexHandler`

- **Login**
  - **URL:** `http://localhost:8100/login`
  - **Descripción:** Página de inicio de sesión.
  - **Método:** `GET, POST`
  - **Handler:** `LoginHandler`

- **Signup**
  - **URL:** `http://localhost:8100/signup`
  - **Descripción:** Página de registro de nuevos usuarios.
  - **Método:** `GET, POST`
  - **Handler:** `SignupHandler`

- **Listado de Autores**
  - **URL:** `http://localhost:8100/authors`
  - **Descripción:** Muestra un listado de todos los autores registrados en la biblioteca.
  - **Método:** `GET`
  - **Handler:** `AuthorsHandler`

- **Formulario para Agregar un Nuevo Autor**
  - **URL:** `http://localhost:8100/author/new`
  - **Descripción:** Formulario donde se puede ingresar información para agregar un nuevo autor a la biblioteca.
  - **Método:** `GET, POST`
  - **Handler:** `NewAuthorHandler`

- **Editar un Autor**
  - **URL:** `http://localhost:8100/author/edit/{id}`
  - **Descripción:** Formulario para editar la información de un autor existente.
  - **Método:** `GET, POST`
  - **Handler:** `EditAuthorHandler`

- **Eliminar un Autor**
  - **URL:** `http://localhost:8100/author/delete/{id}`
  - **Descripción:** Endpoint para eliminar un autor.
  - **Método:** `GET`
  - **Handler:** `DeleteAuthorHandler`

- **Listado de Libros**
  - **URL:** `http://localhost:8100/books`
  - **Descripción:** Muestra un listado de todos los libros registrados en la biblioteca.
  - **Método:** `GET`
  - **Handler:** `BooksHandler`

- **Formulario para Agregar un Nuevo Libro**
  - **URL:** `http://localhost:8100/book/new`
  - **Descripción:** Formulario donde se puede ingresar información para agregar un nuevo libro a la biblioteca.
  - **Método:** `GET, POST`
  - **Handler:** `SaveBookHandler`

- **Editar un Libro**
  - **URL:** `http://localhost:8100/book/edit/{id}`
  - **Descripción:** Formulario para editar la información de un libro existente.
  - **Método:** `GET, POST`
  - **Handler:** `EditBookHandler`

- **Error Handling**
  - **URL:** `http://localhost:8100/error`
  - **Descripción:** Página de error general.
  - **Método:** `GET`
  - **Handler:** `ErrorHandler`

- **Not Found**
  - **URL:** `http://localhost:8100/not-found`
  - **Descripción:** Página de recurso no encontrado.
  - **Método:** `GET`
  - **Handler:** `NotFoundHandler`

## Instrucciones de Uso

1. **Inicio del Servidor**
   - Ejecuta el servidor utilizando el siguiente comando en la raíz del proyecto:
     ```sh
     go run main.go
     ```
   - El servidor estará disponible en `http://localhost:8100`.

2. **Acceso a las Funcionalidades**
   - Abre un navegador web y visita las siguientes URLs para acceder a las diferentes funcionalidades:
     - `http://localhost:8100/` - Página principal con enlaces a autores y libros.
     - `http://localhost:8100/authors` - Listado de autores.
     - `http://localhost:8100/author/new` - Formulario para agregar un nuevo autor.
     - `http://localhost:8100/books` - Listado de libros.
     - `http://localhost:8100/book/new` - Formulario para agregar un nuevo libro.

3. **Interacción con la Aplicación**
   - En las páginas de listado (`/authors` y `/books`), se pueden visualizar los registros actuales.
   - En las páginas de formulario (`/author/new` y `/book/new`), se pueden ingresar nuevos datos y guardarlos en la base de datos.

## Tecnologías Utilizadas

- **Backend:**
  - Go (Golang)
  - GORM (biblioteca de ORM para Go)
  - PostgreSQL (base de datos relacional)

- **Frontend:**
  - HTML
  - CSS (estilos básicos)

## Notas Adicionales

- Asegúrate de tener PostgreSQL instalado y configurado con una base de datos llamada `Biblioteca`.
- Los archivos de plantillas HTML se encuentran en la carpeta `templates/`.
- Los estilos CSS se encuentran en la carpeta `static/`.

## Código del Servidor

```go
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

    // Rutas para el control de errores y métodos no permitidos en la aplicación
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
    log.Printf("Server listening at http://localhost:%d\n", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}