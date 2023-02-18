package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RaoulDucke/product-api/internal/db"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Handler struct {
	repo *db.Repository
}

func New(repository *db.Repository) *Handler {
	return &Handler{repo: repository}
}

func (h *Handler) AddProduct(ctx context.Context, c *gin.Context) {
	req := new(AddProductRequest)
	err := c.BindJSON(req)
	if err != nil {
		internalError(c, err)
		return
	}
	if req.Title == "" {
		badRequest(c)
		return
	}
	if req.Description == "" {
		badRequest(c)
		return
	}
	err = h.repo.AddProduct(ctx, req.Title, req.Description)
	if err != nil {
		internalError(c, err)
		return
	}
}

func (h *Handler) AddProductItem(ctx context.Context, c *gin.Context) {
	req := new(AddProductItemRequest)
	err := c.BindJSON(req)
	if err != nil {
		internalError(c, err)
		return
	}
	if req.Material == "" {
		badRequest(c)
		return
	}
	sku := uuid.NewV4()

	err = h.repo.AddProductItem(ctx, sku.String(), req.Material, req.ProductID)
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

// func notFound(c *gin.Context) {
// 	c.JSON(http.StatusNotFound, "not found")
// }

// func statusOk(c *gin.Context, val any) {
// 	c.JSON(http.StatusOK, val)
// }
