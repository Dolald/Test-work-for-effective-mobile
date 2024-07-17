package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repo interface {
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}
