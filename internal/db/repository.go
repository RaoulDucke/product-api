package db

import (
	"database/sql"
	"errors"

	"golang.org/x/net/context"
)

type Repository struct {
	products []*Product

	database *sql.DB
}

func New(database *sql.DB) *Repository {
	return &Repository{
		products: []*Product{},

		database: database,
	}

}
func (r *Repository) AddProduct(ctx context.Context, title string, description string) error {
	if title == "" {
		return errors.New("title is empty")
	}
	if description == "" {
		return errors.New("description is empty")
	}

	_, err := r.database.ExecContext(ctx, `
			insert into product (title, description)
			values ($1,$2)
		`, title, description)
	if err != nil {
		return err
	}
	return nil
}
