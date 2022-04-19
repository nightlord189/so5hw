package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nightlord189/so5hw/internal/model"
	"net/http"
)

// GetCustomer godoc
// @Tags customer
// @Accept  json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "customer's ID"
// @Success 200 {object} model.CustomerDB
// @Failure 401 {object} model.GenericResponse
// @Failure 422 {object} model.GenericResponse
// @Failure 400 {object} model.GenericResponse
// @Router /customer/{id} [Get]
// @BasePath /
func (h *Handler) GetCustomer(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, model.GenericError(-12, "empty id"))
		return
	}
	entity, err := h.DB.GetCustomer(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.GenericError(-13, "error search: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity)
}
