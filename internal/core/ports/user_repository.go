package ports

import "crud-stefanini/internal/core/domain"

type UserRepository interface {
	Create(user *domain.User) error

	GetByID(id int) (*domain.User, error)

	Update(user *domain.User) error

	GetAll(limit int, offset int) ([]*domain.User, error)
	
	Delete(id int) error
}