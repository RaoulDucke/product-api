package api

type Product struct {
	Identity int64   `json:"id"`
	Name     string  `json:"title"`
	Price    float64 `json:"price"`
}
