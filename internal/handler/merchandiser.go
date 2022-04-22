package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nightlord189/so5hw/internal/model"
	"net/http"
)

// GetMerchandiser godoc
// @Tags merchandiser
// @Accept  json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "merchandiser's ID"
// @Success 200 {object} model.MerchandiserDB
// @Failure 401 {object} model.GenericResponse
// @Failure 422 {object} model.GenericResponse
// @Failure 400 {object} model.GenericResponse
// @Router /api/merchandiser/{id} [Get]
// @BasePath /
func (h *Handler) GetMerchandiser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, model.GenericError(-12, "empty id"))
		return
	}
	var entity model.MerchandiserDB
	err := h.DB.GetEntityByField("id", id, &entity)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.GenericError(-13, "error search: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity)
}
