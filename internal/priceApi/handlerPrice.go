package priceapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RaoulDucke/product-api/internal/db"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *db.Repository
}

func New(repository *db.Repository) *Handler {
	return &Handler{repo: repository}
}
func (h *Handler) AddProductPrice(ctx context.Context, c *gin.Context) {
	req := new(AddProductPriceRequest)
	err := c.BindJSON(req)
	if err != nil {
		internalError(c, err)
		return
	}
	if req.Price <= 0 {
		badRequest(c)
		return
	}

	err = h.repo.AddProductPrice(ctx, req.ProductID, req.Price)
	if err != nil {
		_, ok := err.(*db.ErrNotFound)
		if ok {
			badRequest(c)
			return
		}
		// if err == db.ErrProductNotFound {
		// 	badRequest(c)
		// }
		internalError(c, err)
		return
	}
}

func internalError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, fmt.Sprintf("internal error: %s", err))
}

func badRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, "bad request")
}
