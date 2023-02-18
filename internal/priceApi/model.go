package priceapi

type AddProductPriceRequest struct {
	ProductID int `json: "productID"`
	Price     int `json: "price"`
}
