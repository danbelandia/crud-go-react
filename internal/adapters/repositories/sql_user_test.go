package repositories

import (
	"database/sql"
	"testing"

	"crud-stefanini/internal/core/domain"
	_ "modernc.org/sqlite"
)

// Las funciones de prueba en Go SIEMPRE empiezan con "Test" y reciben (t *testing.T)
func TestSQLiteUserRepository_Create(t *testing.T) {
	// 1. Preparación (Setup): Levantamos una base de datos EN MEMORIA
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Error abriendo BD en memoria: %v", err)
	}
	defer db.Close()

	// Creamos la tabla solo para esta prueba
	db.Exec(`CREATE TABLE users (id integer PRIMARY KEY AUTOINCREMENT, name TEXT, last_name TEXT, email TEXT);`)

	// Instanciamos nuestro repositorio real
	repo := NewSqlUserRepository(db)

	// 2. Ejecución (Act)
	nuevoUser := &domain.User{
		Name:     "Test Name",
		LastName: "Test Lastname",
		Email:    "test@ejemplo.com",
	}
	err = repo.Create(nuevoUser)

	// 3. Verificación (Assert)
	if err != nil {
		// t.Errorf marca la prueba como fallida pero sigue ejecutando
		t.Errorf("Se esperaba error nulo, se obtuvo: %v", err) 
	}

	if nuevoUser.ID == 0 {
		t.Errorf("Se esperaba que SQLite asignara un ID mayor a 0, pero es %d", nuevoUser.ID)
	}
}