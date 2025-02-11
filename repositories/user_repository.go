package repositories

import (
	"database/sql"

	"gin-simple.com/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	// Implementation
}

func (r *UserRepository) GetByID(id int64) (*models.User, error) {
	// Implementation
}

func (r *UserRepository) Create(user *models.User) error {
	// Implementation
}

func (r *UserRepository) Update(user *models.User) error {
	// Implementation
}

func (r *UserRepository) Delete(id int64) error {
	// Implementation
}
