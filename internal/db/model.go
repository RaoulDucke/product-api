package db

type Product struct {
	ID          int64
	Title       string
	Description string
}

type productItem struct {
	ID        int
	Sku       string
	Material  string
	ProductID int
}

type ProductPrice struct {
	ID        int
	ProductID int
	Price     int
}

// var ErrProductNotFound = errors.New("product not found")

type ErrNotFound struct {
	massange string
}

func (e *ErrNotFound) Error() string {
	return e.massange
}
