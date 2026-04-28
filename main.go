package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"crud-stefanini/internal/adapters/handlers"
	"crud-stefanini/internal/adapters/repositories"
	"crud-stefanini/internal/utils"

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
	//Agregamos el logger a los handlers
	handlersConLogger := utils.Logger(userHandler.HandleUsuarios)
	handlerCompleto := CorsMiddleware(handlersConLogger)
	//Rutas usuarios y usuarios/id
	http.HandleFunc("/usuarios", handlerCompleto)
	http.HandleFunc("/usuarios/", handlerCompleto)

	fmt.Println("Servidor -> http://localhost:8080/usuarios")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
