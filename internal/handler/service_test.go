package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nightlord189/so5hw/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestResetDB(t *testing.T) {
	customers := []model.CustomerDB{
		{
			Email:           "test35@mail.com",
			BillingAddress:  "some city, some street",
			ShippingAddress: "different city, new street",
		},
	}
	err := dbInstance.CreateEntities(&customers)
	require.NoError(t, err)

	merchandiser := model.MerchandiserDB{
		Username: "merchandiser01",
	}
	err = dbInstance.CreateEntity(&merchandiser)
	require.NoError(t, err)

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		handlerInstance.ResetDB(ct)
		p, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, http.StatusOK, w.Code, string(p))
		assert.True(t, strings.Contains(string(p), "success"))

		var customer model.CustomerDB
		err = dbInstance.GetEntityByField("email", customers[0].Email, &customer)
		assert.Error(t, err)
	})
}

func TestFillDB(t *testing.T) {
	err := dbInstance.TruncateAllTables()
	require.NoError(t, err)

	t.Cleanup(func() {
		err := dbInstance.TruncateAllTables()
		assert.NoError(t, err)
	})

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		handlerInstance.FillDB(ct)
		p, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, http.StatusOK, w.Code, string(p))
		assert.True(t, strings.Contains(string(p), "success"))
	})
}
