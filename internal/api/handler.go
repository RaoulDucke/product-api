package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	r *repository.Repository
}

func New(repository *repository.Repository) *Handler {
	return &Handler{r: repository}
}

func (h *Handler) AddProduct(ctx context.Context, c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		internalError(c, err)
		return
	}
	product := new(Product)
	err = json.Unmarshal(jsonData, product)
	if err != nil {
		internalError(c, err)
		return
	}
	if product.Name == "" {
		badRequst(c)
		return
	}
	if product.Price <= 0 {
		badRequst(c)
		return
	}
	err = h.r.AddProduct(ctx, convertToDBProduct(product))
	if err != nil {
		internalError(c, err)
		return
	}
}

func (h *Handler) GetProducts(ctx context.Context, c *gin.Context) {
	idString := c.Request.URL.Query().Get("id")
	if idString != "" {
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			badRequst(c)
			return
		}
		product, ok := h.getProduct(id)
		if ok {
			statusOk(c, product)
		} else {
			notFound(c)
		}
		return
	}
	products, err := h.r.GetProducts(ctx)
	if err != nil {
		internalError(c, err)
		return
	}

	c.JSON(http.StatusOK, convertToProducts(products))

}

func convertToDBProduct(p *Product) *repository.Product {
	return &repository.Product{
		Title: p.Name,
		ID:    p.Identity,
		Price: p.Price,
	}
}

func internalError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, fmt.Sprintf("internal error: %s", err))
}

func badRequst(c *gin.Context) {
	c.JSON(http.StatusBadRequest, "bad request")
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, "not found")
}

func statusOk(c *gin.Context, val any) {
	c.JSON(http.StatusOK, val)
}
