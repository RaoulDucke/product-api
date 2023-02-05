package db

import (
	"context"
	"database/sql"
	"errors"
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
func (r *Repository) AddProduct(ctx context.Context, p *Product) error {
	if p == nil {
		return errors.New("product is nil")
	}
	if p.Title == "" {
		return errors.New("title is empty")
	}
	if p.Price <= 0 {
		return errors.New("price <=0")
	}
	id := int64(1)
	if len(r.products) > 0 {
		lastProduct := r.products[len(r.products)-1]
		id = lastProduct.ID + 1
	}

	p.ID = id
	r.products = append(r.products, p)
	return nil

}

func (r *Repository) GetProducts(ctx context.Context) ([]*Product, error) {
	raws, err := r.database.QueryContext(ctx, `
	select id,title, price
	from product
	`)
	if err != nil {
		return nil, err
	}

	defer raws.Close()
	var result []*Product

	for raws.Next() {
		p := new(Product)
		err = raws.Scan(&p.ID, &p.Title, &p.Price)
		if err != nil {
			return nil, err
		}

		result = append(result, p)
	}

	return result, nil
}
