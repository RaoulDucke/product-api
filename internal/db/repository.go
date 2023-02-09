package db

import (
	"database/sql"
	"errors"

	"golang.org/x/net/context"
)

type Repository struct {
	database *sql.DB
}

func New(database *sql.DB) *Repository {
	return &Repository{
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

func (r *Repository) AddProductItem(ctx context.Context, sku string, material string, productID int) error {
	if material == "" {
		return errors.New("material is empty")
	}
	_, err := r.database.ExecContext(ctx, `
			insert into product_item (sku, material, product_id)
			values ($1,$2,$3)
		`, sku, material, productID)
	if err != nil {
		return err
	}
	return nil
}
