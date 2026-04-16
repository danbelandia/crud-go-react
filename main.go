package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"crud-stefanini/internal/adapters/handlers"
	"crud-stefanini/internal/adapters/repositories"

	_ "modernc.org/sqlite"
)

func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permite a cualquier frontend
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Role")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func main() {
	//Instanciamos la BD
	db, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec(`CREATE TABLE IF NOT EXISTS users (id integer PRIMARY KEY AUTOINCREMENT, name TEXT, last_name TEXT, email TEXT, is_admin BOOLEAN DEFAULT 0);`)
	//Instanciamos el repositorio
	repo := repositories.NewSqlUserRepository(db)
	//Capa de HTTP
	userHandler := handlers.NewUserHandler(repo)
	//Con esto evitamos el cors
	handlerConCors := CorsMiddleware(userHandler.HandleUsuarios)
	//Ruta y usamos función principal del handler
	http.HandleFunc("/usuarios", handlerConCors)
	http.HandleFunc("/usuarios/", handlerConCors)

	fmt.Println("Servidor -> http://localhost:8080/usuarios")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
