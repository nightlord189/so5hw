package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nightlord189/so5hw/internal/model"
	"github.com/nightlord189/so5hw/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestAuth(t *testing.T) {
	password := "good_password"
	customer := model.CustomerDB{
		Email:        "test@mail.com",
		PasswordHash: service.HashPassword(password),
	}
	err := dbInstance.CreateEntity(&customer)
	require.NoError(t, err)

	merchandiser := model.MerchandiserDB{
		Username:     "merchandiser01",
		PasswordHash: service.HashPassword(password),
	}
	err = dbInstance.CreateEntity(&merchandiser)
	require.NoError(t, err)

	t.Run("Non-existent user", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		payloadStruct := model.AuthRequest{
			Username: "non_existent_user",
			Password: "123456",
			Type:     model.UserTypeCustomer,
		}
		payload, _ := json.Marshal(payloadStruct)
		ct.Request = &http.Request{
			Body: ioutil.NopCloser(bytes.NewBuffer(payload)),
		}
		handlerInstance.Auth(ct)
		p, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.True(t, strings.Contains(string(p), "bad credentials"))
	})

	t.Run("Wrong password", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		payloadStruct := model.AuthRequest{
			Username: customer.Email,
			Password: "wrong_password",
			Type:     model.UserTypeCustomer,
		}
		payload, _ := json.Marshal(payloadStruct)
		ct.Request = &http.Request{
			Body: ioutil.NopCloser(bytes.NewBuffer(payload)),
		}
		handlerInstance.Auth(ct)
		p, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.True(t, strings.Contains(string(p), "bad credentials"))
	})

	t.Run("Auth customer", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		payloadStruct := model.AuthRequest{
			Username: customer.Email,
			Password: password,
			Type:     model.UserTypeCustomer,
		}
		payload, _ := json.Marshal(payloadStruct)
		ct.Request = &http.Request{
			Body: ioutil.NopCloser(bytes.NewBuffer(payload)),
		}
		handlerInstance.Auth(ct)

		p, _ := ioutil.ReadAll(w.Body)

		assert.Equal(t, http.StatusOK, w.Code)

		claim, err := service.ValidateJwtToken(string(p), configInstance.AuthAccessSecret)
		require.NoError(t, err)
		assert.EqualValues(t, customer.ID, claim["user_id"])
		assert.EqualValues(t, customer.Email, claim["username"])
		assert.EqualValues(t, model.UserTypeCustomer, claim["role"])
		assert.GreaterOrEqual(t, claim["exp"].(float64),
			float64(time.Now().Add(time.Second*time.Duration(handlerInstance.Config.TokenExpTime)).Unix()))
	})

	t.Run("Auth merchandiser", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		payloadStruct := model.AuthRequest{
			Username: merchandiser.Username,
			Password: password,
			Type:     model.UserTypeMerchandiser,
		}
		payload, _ := json.Marshal(payloadStruct)
		ct.Request = &http.Request{
			Body: ioutil.NopCloser(bytes.NewBuffer(payload)),
		}
		handlerInstance.Auth(ct)

		p, _ := ioutil.ReadAll(w.Body)

		assert.Equal(t, http.StatusOK, w.Code)

		claim, err := service.ValidateJwtToken(string(p), configInstance.AuthAccessSecret)
		require.NoError(t, err)
		assert.EqualValues(t, merchandiser.ID, claim["user_id"])
		assert.EqualValues(t, merchandiser.Username, claim["username"])
		assert.EqualValues(t, model.UserTypeMerchandiser, claim["role"])
		assert.GreaterOrEqual(t, claim["exp"].(float64),
			float64(time.Now().Add(time.Second*time.Duration(handlerInstance.Config.TokenExpTime)).Unix()))
	})
}
