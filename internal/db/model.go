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
