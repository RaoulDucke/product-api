package db

import (
	"database/sql"
	"errors"

	// "net/http"

	// "github.com/gin-gonic/gin"
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
	// 	var res int64
	// 	err := r.database.QueryRowContext(ctx, "SELECT id FROMproduct WHERE id = $1", productID).Scan(&res)
	// 	if err != nil {
	// 		if err == sql.ErrNoRows {
	// 			return ErrProductNotFound
	// 		}
	// 		return err
	// 	}
	var res int64
	err := r.database.QueryRowContext(ctx, "SELECT id FROMproduct WHERE id = $1", productID).Scan(&res)
	if err != nil {
		if err == sql.ErrNoRows {
			return &ErrNotFound{
				massange: "product not found",
			}
		}
		return err
	}

	_, err = r.database.ExecContext(ctx, `
			insert into product_item (sku, material, product_id)
			values ($1,$2,$3)
		`, sku, material, productID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) AddProductPrice(ctx context.Context, productID int, price int) error {
	if price <= 0 {
		return &ErrNotFound{
			massange: "price <= 0",
		}
	}
	var req int
	err := r.database.QueryRowContext(ctx, "SELECT id FROM product WHERE id = $1", productID).Scan(&req)
	if err != nil {
		if err == sql.ErrNoRows {
			return &ErrNotFound{
				massange: "product not found",
			}
		}
		return err
	}
	_, err = r.database.ExecContext(ctx, `
			insert into product_price (product_id, price)
			values ($1,$2)
		`, productID, price)
	if err != nil {
		return err
	}
	return nil
}

// func badRequest(c *gin.Context) {
// 	c.JSON(http.StatusBadRequest, "bad request")
// }
