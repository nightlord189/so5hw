package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nightlord189/so5hw/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestSale(t *testing.T) {
	creditCard := model.CreditCardInfo{
		CVV:    "591",
		Number: "4024007145804521",
		Holder: "John Beaver",
		Date:   "11/26",
	}
	creditCardJson, err := json.Marshal(creditCard)
	assert.NoError(t, err)
	customers := []model.CustomerDB{
		{
			Email:           "test@mail.com",
			BillingAddress:  "some city, some street",
			ShippingAddress: "different city, new street",
			CreditCardRaw:   string(creditCardJson),
		},
		{
			Email:           "test2@mail.com",
			BillingAddress:  "some city, some street",
			ShippingAddress: "different city, new street",
		},
	}
	err = dbInstance.CreateEntities(&customers)
	require.NoError(t, err)

	products := []model.ProductDB{
		{
			Articul:                 "art111",
			Price:                   20,
			DeliveryTimeDescription: "2-5 days",
			Category:                "category1",
			Status:                  model.ProductStatusActive,
			Inventory:               10,
			Vendor:                  "vendor1",
		},
		{
			Articul:                 "art112",
			Price:                   25,
			DeliveryTimeDescription: "3-5 days",
			Category:                "category1",
			Status:                  model.ProductStatusInactive,
			Inventory:               1,
			Vendor:                  "vendor1",
		},
	}
	err = dbInstance.CreateEntities(&products)
	require.NoError(t, err)

	t.Run("Empty body", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		handlerInstance.Sale(ct)
		p, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.True(t, strings.Contains(string(p), "error bind model"))
	})

	t.Run("Non-existent customer", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		payloadStruct := model.SaleRequest{
			CustomerID: 404,
			ProductID:  products[0].ID,
			Quantity:   1,
		}
		payload, _ := json.Marshal(payloadStruct)
		ct.Request = &http.Request{
			Body: ioutil.NopCloser(bytes.NewBuffer(payload)),
		}
		ct.Set("user_id", "404")
		handlerInstance.Sale(ct)
		p, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.True(t, strings.Contains(string(p), "error search customer"))
	})

	t.Run("Wrong user_id", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		payloadStruct := model.SaleRequest{
			CustomerID: customers[0].ID,
			ProductID:  products[0].ID,
			Quantity:   1,
		}
		payload, _ := json.Marshal(payloadStruct)
		ct.Request = &http.Request{
			Body: ioutil.NopCloser(bytes.NewBuffer(payload)),
		}
		ct.Set("user_id", "404")
		handlerInstance.Sale(ct)
		p, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.True(t, strings.Contains(string(p), "wrong customer id in request"))
	})

	t.Run("Non-existent product", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		payloadStruct := model.SaleRequest{
			CustomerID: customers[0].ID,
			ProductID:  404,
			Quantity:   1,
		}
		payload, _ := json.Marshal(payloadStruct)
		ct.Request = &http.Request{
			Body: ioutil.NopCloser(bytes.NewBuffer(payload)),
		}
		ct.Set("user_id", strconv.Itoa(customers[0].ID))
		handlerInstance.Sale(ct)
		p, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.True(t, strings.Contains(string(p), "error search product"))
	})

	t.Run("Inactive product", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		payloadStruct := model.SaleRequest{
			CustomerID: customers[0].ID,
			ProductID:  products[1].ID,
			Quantity:   1,
		}
		payload, _ := json.Marshal(payloadStruct)
		ct.Request = &http.Request{
			Body: ioutil.NopCloser(bytes.NewBuffer(payload)),
		}
		ct.Set("user_id", strconv.Itoa(customers[0].ID))
		handlerInstance.Sale(ct)
		p, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.True(t, strings.Contains(string(p), "product status is not active"))
	})

	t.Run("Not enough inventory", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		payloadStruct := model.SaleRequest{
			CustomerID: customers[0].ID,
			ProductID:  products[0].ID,
			Quantity:   100,
		}
		payload, _ := json.Marshal(payloadStruct)
		ct.Request = &http.Request{
			Body: ioutil.NopCloser(bytes.NewBuffer(payload)),
		}
		ct.Set("user_id", strconv.Itoa(customers[0].ID))
		handlerInstance.Sale(ct)
		p, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.True(t, strings.Contains(string(p), "not enough inventory"))
	})

	t.Run("Customer with empty payment data", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		payloadStruct := model.SaleRequest{
			CustomerID: customers[1].ID,
			ProductID:  products[0].ID,
			Quantity:   2,
		}
		payload, _ := json.Marshal(payloadStruct)
		ct.Request = &http.Request{
			Body: ioutil.NopCloser(bytes.NewBuffer(payload)),
		}
		ct.Set("user_id", strconv.Itoa(customers[1].ID))
		handlerInstance.Sale(ct)
		p, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code, string(p))
		assert.True(t, strings.Contains(string(p), "customer doesn't have filled data for 1-click payment"), string(p))
	})

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		payloadStruct := model.SaleRequest{
			CustomerID: customers[0].ID,
			ProductID:  products[0].ID,
			Quantity:   2,
		}
		payload, _ := json.Marshal(payloadStruct)
		ct.Request = &http.Request{
			Body: ioutil.NopCloser(bytes.NewBuffer(payload)),
		}
		ct.Set("user_id", strconv.Itoa(customers[0].ID))
		handlerInstance.Sale(ct)
		p, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, http.StatusOK, w.Code, string(p))
		assert.True(t, strings.Contains(string(p), "success"))
	})
}
