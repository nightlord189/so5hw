package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nightlord189/so5hw/internal/model"
	"net/http"
	"strconv"
)

// Sale godoc
// @Tags sale
// @Accept  json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param model body model.SaleRequest true "input model"
// @Success 200 {object} model.GenericResponse
// @Failure 401 {object} model.GenericResponse
// @Failure 403 {object} model.GenericResponse
// @Failure 422 {object} model.GenericResponse
// @Failure 400 {object} model.GenericResponse
// @Router /api/sale [Post]
// @BasePath /
func (h *Handler) Sale(c *gin.Context) {
	var req model.SaleRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GenericError(-12, fmt.Sprintf("error bind model: %v", err)))
		return
	}

	customerIDStr := strconv.Itoa(req.CustomerID)
	customer, err := h.DB.GetCustomer(customerIDStr)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.GenericError(-13, "error search customer: "+err.Error()))
		return
	}
	userIdFromClaims := c.GetString("user_id")
	if userIdFromClaims != customerIDStr {
		c.JSON(http.StatusForbidden, model.GenericError(-14, "wrong customer id in request"))
		return
	}

	var product model.ProductDB
	err = h.DB.GetEntityByField("id", fmt.Sprintf("%d", req.ProductID), &product)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.GenericError(-15, "error search product: "+err.Error()))
		return
	}

	if product.Status != model.ProductStatusActive {
		c.JSON(http.StatusUnprocessableEntity, model.GenericError(-16, "product status is not active"))
		return
	}

	if product.Inventory < req.Quantity {
		c.JSON(http.StatusUnprocessableEntity, model.GenericError(-17, "not enough inventory"))
		return
	}

	if !customer.IsFilled() {
		c.JSON(http.StatusUnprocessableEntity,
			model.GenericError(-18, "customer doesn't have filled data for 1-click payment"))
		return
	}

	product.Inventory -= req.Quantity
	err = h.DB.UpdateEntity(&product)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.GenericError(-19, "error update db: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.GenericError(0, "success"))
}
