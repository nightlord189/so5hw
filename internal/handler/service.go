package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nightlord189/so5hw/internal/model"
	"net/http"
)

// ResetDB godoc
// @Tags service
// @Accept  json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} model.GenericResponse
// @Failure 422 {object} model.GenericResponse
// @Router /reset [Post]
// @BasePath /
func (h *Handler) ResetDB(c *gin.Context) {
	err := h.DB.TruncateAllTables()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.GenericError(-11, "error clear db: "+err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.GenericError(0, "success"))
}

// FillDB godoc
// @Tags service
// @Accept  json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} model.GenericResponse
// @Failure 422 {object} model.GenericResponse
// @Router /fill [Post]
// @BasePath /
func (h *Handler) FillDB(c *gin.Context) {
	err := h.DB.FillData()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.GenericError(-11, "error fill db: "+err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.GenericError(0, "success"))
}
