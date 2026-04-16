package repositories

import (
	"crud-stefanini/internal/core/domain"
	"database/sql"
	"errors"
)

type SQLiteUserRepository struct {
	db *sql.DB
}

func NewSqlUserRepository(db *sql.DB) *SQLiteUserRepository {
	return &SQLiteUserRepository{db: db}
}

// Crear usuario
func (r *SQLiteUserRepository) Create(user *domain.User) error {

	query := `INSERT INTO users (name, last_name, email, is_admin) VALUES (?, ?, ?, ?)`

	result, err := r.db.Exec(query, user.Name, user.LastName, user.Email, user.IsAdmin)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}

// Obtner un usuario por ID
func (r *SQLiteUserRepository) GetByID(id int) (*domain.User, error) {
	query := `SELECT id, name, last_name, email, is_admin FROM users WHERE id = ?`

	user := &domain.User{}

	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.IsAdmin)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no se encontró usuario")
		}
		return nil, err
	}
	return user, nil
}

// Actualizar un usuario
func (r *SQLiteUserRepository) Update(user *domain.User) error {
	query := `UPDATE users SET name = ?, last_name = ?, email = ?, is_admin = ? WHERE id = ?`

	result, err := r.db.Exec(query, user.Name, user.LastName, user.Email, user.IsAdmin, user.ID)
	if err != nil {
		return err
	}

	filasAfectadas, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if filasAfectadas == 0 {
		return errors.New("no existe el usuario para actualizar")
	}
	return nil
}

// Eliminar un usuario
func (r *SQLiteUserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	filasAfectadas, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if filasAfectadas == 0 {
		return errors.New("no se encontró ningún usuario para eliminar")
	}
	return nil
}

// Obtener todos los usuarios
func (r *SQLiteUserRepository) GetAll(limit int, offset int) ([]*domain.User, error) {
	query := `SELECT id, name, last_name, email, is_admin FROM users LIMIT ? OFFSET ?`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close() //Cierra las filas

	var users []*domain.User

	for rows.Next() {
		u := &domain.User{}
		err := rows.Scan(&u.ID, &u.Name, &u.LastName, &u.Email, &u.IsAdmin)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}
	return users, nil
}
