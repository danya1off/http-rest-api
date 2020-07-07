package user

import (
	"github.com/danya1off/http-rest-api/internal/app/store"
	"github.com/sirupsen/logrus"
)

// Repository ...
type Repository struct {
	DB     store.DB
	Logger *logrus.Logger
}

// Create ...
func (r Repository) Create(u *User) (*User, error) {
	r.Logger.Info("Create was called")
	return nil, nil
}

// FindByEmail ...
func (r Repository) FindByEmail(email string) (*User, error) {
	r.Logger.Info("Find by email was called")
	return nil, nil
}
