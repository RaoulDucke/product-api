package api

type AddProductRequest struct {
	Title       string `json:"title"`
	Description string `json: "description"`
}

type AddProductItemRequest struct {
	Material  string `json: "material"`
	ProductID int    `json: "productID"`
}
