package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/nightlord189/so5hw/internal/model"
	"github.com/nightlord189/so5hw/internal/service"
	"net/http"
	"time"
)

//Auth godoc
//@Summary Auth request
//@Description Request to authorize via JWT-token
//@Tags auth
//@Accept  json
//@Produce json
//@Param data body model.AuthRequest true "Input model"
//@Success 200 {string} string
//@Failure 401 {object} model.GenericResponse
//@Failure 422 {object} model.GenericResponse
//@Failure 400 {object} model.GenericResponse
//@Router /api/auth [Post]
//@BasePath /
func (h *Handler) Auth(c *gin.Context) {
	var req model.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.GenericError(-11, fmt.Sprintf("error parse json: %v", err)))
		return
	}
	user, err := h.DB.GetUserEntity(req.Username, req.Type)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.GenericError(-31, "bad credentials"))
		return
	}
	if service.HashPassword(req.Password) != user.PasswordHash {
		c.JSON(http.StatusUnauthorized, model.GenericError(-31, "bad credentials"))
		return
	}
	payload := jwt.MapClaims{}
	payload["user_id"] = user.ID
	payload["username"] = req.Username
	payload["role"] = req.Type
	payload["exp"] = time.Now().Add(time.Second * time.Duration(h.Config.TokenExpTime)).Unix()
	payload["iss"] = "so5hw"
	token, err := service.CreateToken(payload, h.Config.AuthAccessSecret)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.GenericError(-41, "error create token"))
		return
	}
	c.String(http.StatusOK, token)
}
