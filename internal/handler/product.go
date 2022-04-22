package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nightlord189/so5hw/internal/model"
	"net/http"
)

// GetCategories godoc
// @Tags product
// @Accept  json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Success 200 {array} string
// @Failure 401 {object} model.GenericResponse
// @Failure 422 {object} model.GenericResponse
// @Router /api/product/category [Get]
// @BasePath /
func (h *Handler) GetCategories(c *gin.Context) {
	response, err := h.DB.GetCategories()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.GenericError(-14, fmt.Sprintf("err get: %v", err)))
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetProducts godoc
// @Tags product
// @Accept  json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param id query string false "product's ID"
// @Param articul query string false "articul"
// @Param category query string false "category"
// @Param status query model.ProductStatus false "status"
// @Param vendor query string false "vendor"
// @Param limit query integer false "Page's size"
// @Param page query integer false "Number of page"
// @Success 200 {object} model.GetProductsRequest
// @Failure 401 {object} model.GenericResponse
// @Failure 422 {object} model.GenericResponse
// @Failure 400 {object} model.GenericResponse
// @Router /api/product [Get]
// @BasePath /
func (h *Handler) GetProducts(c *gin.Context) {
	var req model.GetProductsRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GenericError(-12, fmt.Sprintf("err binding query: %v", err)))
		return
	}

	response, err := h.DB.GetProducts(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.GenericError(-14, fmt.Sprintf("err get entities: %v", err)))
		return
	}
	c.JSON(http.StatusOK, response)
}

// CreateProduct godoc
// @Tags product
// @Accept  json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param model body model.CreateProductRequest true "input model"
// @Success 200 {object} model.ProductDB
// @Failure 401 {object} model.GenericResponse
// @Failure 422 {object} model.GenericResponse
// @Failure 400 {object} model.GenericResponse
// @Router /api/product [Post]
// @BasePath /
func (h *Handler) CreateProduct(c *gin.Context) {
	var req model.CreateProductRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GenericError(-12, fmt.Sprintf("error bind model: %v", err)))
		return
	}

	product, err := h.DB.CreateProduct(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.GenericError(-17, "error create product: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Tags product
// @Accept  json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "product's ID"
// @Success 200 {object} model.GenericResponse
// @Failure 401 {object} model.GenericResponse
// @Failure 422 {object} model.GenericResponse
// @Failure 400 {object} model.GenericResponse
// @Router /api/product/{id} [Delete]
// @BasePath /
func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, model.GenericError(-12, "empty id"))
		return
	}

	err := h.DB.DeleteEntityByField("id", id, model.ProductDB{})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.GenericError(-15, "error delete: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.GenericError(0, "success"))
}
