package handler

import (
	"bytes"
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
	"net/url"
	"strings"
	"testing"
)

func TestDeleteProduct(t *testing.T) {
	entity := model.ProductDB{
		Articul:                 "art35",
		Price:                   100,
		DeliveryTimeDescription: "1 day",
		Category:                "category1",
		Status:                  model.ProductStatusActive,
		Inventory:               1,
		Vendor:                  "vendor35",
	}
	err := dbInstance.CreateEntity(&entity)
	require.NoError(t, err)

	t.Run("Empty id", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		handlerInstance.DeleteProduct(ct)
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
				Value: fmt.Sprintf("%d", entity.ID),
			},
		}
		handlerInstance.DeleteProduct(ct)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestCreateProduct(t *testing.T) {
	t.Run("Empty body", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		handlerInstance.CreateProduct(ct)
		p, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.True(t, strings.Contains(string(p), "error bind model"))
	})

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		payloadStruct := model.CreateProductRequest{
			Articul:                 "tassay891",
			Price:                   101,
			DeliveryTimeDescription: "1 day",
			Status:                  model.ProductStatusActive,
			Inventory:               1,
			Vendor:                  "vendor801",
			Category:                "water",
			Images:                  [][]byte{{1, 2, 3}, {5, 6, 7}},
		}
		payload, _ := json.Marshal(payloadStruct)
		ct.Request = &http.Request{
			Body: ioutil.NopCloser(bytes.NewBuffer(payload)),
		}

		handlerInstance.CreateProduct(ct)
		p, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, http.StatusOK, w.Code)
		var response model.ProductDB
		err := json.Unmarshal(p, &response)
		require.NoError(t, err)
		assert.Equal(t, payloadStruct.Status, response.Status)
		assert.Equal(t, payloadStruct.Price, response.Price)
		assert.Equal(t, payloadStruct.Category, response.Category)
		assert.Equal(t, payloadStruct.Vendor, response.Vendor)
		assert.Equal(t, payloadStruct.Articul, response.Articul)
		assert.Equal(t, payloadStruct.Inventory, response.Inventory)
		assert.Equal(t, payloadStruct.DeliveryTimeDescription, response.DeliveryTimeDescription)
		assert.Len(t, response.Images, len(payloadStruct.Images))
	})
}

func TestGetCategories(t *testing.T) {
	products := []model.ProductDB{
		{
			Articul:                 "art301",
			Price:                   20,
			DeliveryTimeDescription: "2-5 days",
			Category:                "category1",
			Status:                  model.ProductStatusActive,
			Inventory:               3,
			Vendor:                  "vendor1",
		},
		{
			Articul:                 "art302",
			Price:                   25,
			DeliveryTimeDescription: "3-5 days",
			Category:                "category1",
			Status:                  model.ProductStatusActive,
			Inventory:               1,
			Vendor:                  "vendor1",
		},
		{
			Articul:                 "art303",
			Price:                   30,
			DeliveryTimeDescription: "4-5 days",
			Category:                "category2",
			Status:                  model.ProductStatusInactive,
			Inventory:               0,
			Vendor:                  "vendor1",
		},
		{
			Articul:                 "art304",
			Price:                   30,
			DeliveryTimeDescription: "4-5 days",
			Category:                "water",
			Status:                  model.ProductStatusActive,
			Inventory:               0,
			Vendor:                  "vendor1",
		},
	}
	err := dbInstance.CreateEntities(&products)
	require.NoError(t, err)

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		handlerInstance.GetCategories(ct)
		p, _ := io.ReadAll(w.Body)
		assert.Equal(t, http.StatusOK, w.Code, string(p))
		var response []string
		err := json.Unmarshal(p, &response)
		require.NoError(t, err)
		assert.Len(t, response, 2)
	})
}

func TestGetGoods(t *testing.T) {
	products := []model.ProductDB{
		{
			Articul:                 "art201",
			Price:                   20,
			DeliveryTimeDescription: "2-5 days",
			Category:                "category201",
			Status:                  model.ProductStatusActive,
			Inventory:               3,
			Vendor:                  "vendor201",
		},
		{
			Articul:                 "art202",
			Price:                   25,
			DeliveryTimeDescription: "3-5 days",
			Category:                "category201",
			Status:                  model.ProductStatusActive,
			Inventory:               1,
			Vendor:                  "vendor201",
		},
		{
			Articul:                 "art203",
			Price:                   30,
			DeliveryTimeDescription: "4-5 days",
			Category:                "category202",
			Status:                  model.ProductStatusInactive,
			Inventory:               0,
			Vendor:                  "vendor201",
		},
		{
			Articul:                 "art204",
			Price:                   50,
			DeliveryTimeDescription: "3-5 days",
			Category:                "category202",
			Status:                  model.ProductStatusActive,
			Inventory:               7,
			Vendor:                  "vendor202",
		},
	}
	err := dbInstance.CreateEntities(&products)
	require.NoError(t, err)

	t.Run("Filter by vendor", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		u, err := url.Parse("/?vendor=vendor201&limit=2&page=2")
		assert.NoError(t, err)
		ct.Request = &http.Request{
			URL: u,
		}
		handlerInstance.GetProducts(ct)
		p, _ := io.ReadAll(w.Body)
		assert.Equal(t, http.StatusOK, w.Code, string(p))
		var response model.GetProductsResponse
		err = json.Unmarshal(p, &response)
		require.NoError(t, err)
		assert.Len(t, response.Records, 1)
		assert.Equal(t, products[2], response.Records[0])
		assert.Equal(t, 3, response.RecordsCount)
		assert.Equal(t, 2, response.CurrentPage)
		assert.Equal(t, 2, response.PagesCount)
	})

	t.Run("Filter by status", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		u, err := url.Parse("/?status=active&vendor=vendor201")
		assert.NoError(t, err)
		ct.Request = &http.Request{
			URL: u,
		}
		handlerInstance.GetProducts(ct)
		p, _ := io.ReadAll(w.Body)
		assert.Equal(t, http.StatusOK, w.Code, string(p))
		var response model.GetProductsResponse
		err = json.Unmarshal(p, &response)
		require.NoError(t, err)
		assert.Len(t, response.Records, 2)
	})

	t.Run("Filter by category", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		u, err := url.Parse("/?category=category1")
		assert.NoError(t, err)
		ct.Request = &http.Request{
			URL: u,
		}
		handlerInstance.GetProducts(ct)
		p, _ := io.ReadAll(w.Body)
		assert.Equal(t, http.StatusOK, w.Code, string(p))
		var response model.GetProductsResponse
		err = json.Unmarshal(p, &response)
		require.NoError(t, err)
		assert.Len(t, response.Records, 2)
	})
}
