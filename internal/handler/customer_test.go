package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nightlord189/so5hw/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetCustomer(t *testing.T) {
	creditCard := model.CreditCardInfo{
		CVV:    "591",
		Number: "4024007145804521",
		Holder: "John Beaver",
		Date:   "11/26",
	}
	creditCardJson, err := json.Marshal(creditCard)
	assert.NoError(t, err)
	customer := model.CustomerDB{
		Email:           "test@mail.com",
		BillingAddress:  "some city, some street",
		ShippingAddress: "different city, new street",
		CreditCardRaw:   string(creditCardJson),
	}
	err = dbInstance.CreateEntity(&customer)
	require.NoError(t, err)

	t.Run("Empty id", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		handlerInstance.GetCustomer(ct)
		p, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.True(t, strings.Contains(string(p), "empty id"))
	})

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		ct.Params = []gin.Param{
			{
				Key:   "id",
				Value: fmt.Sprintf("%d", customer.ID),
			},
		}
		handlerInstance.GetCustomer(ct)
		p, _ := io.ReadAll(w.Body)
		assert.Equal(t, http.StatusOK, w.Code)
		var response model.CustomerDB
		err = json.Unmarshal(p, &response)
		require.NoError(t, err)
		assert.Equal(t, customer.ID, response.ID)
		assert.Equal(t, customer.Email, response.Email)
		assert.Equal(t, customer.ShippingAddress, response.ShippingAddress)
		assert.Equal(t, customer.BillingAddress, response.BillingAddress)
		assert.Equal(t, creditCard, response.CreditCard)
	})
}
