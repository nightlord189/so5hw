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

func TestGetMerchandiser(t *testing.T) {
	merchandiser := model.MerchandiserDB{
		Username: "merchandiser01",
	}
	err := dbInstance.CreateEntity(&merchandiser)
	require.NoError(t, err)

	t.Run("Empty id", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		ct, _ := gin.CreateTestContext(w)
		handlerInstance.GetMerchandiser(ct)
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
				Value: fmt.Sprintf("%d", merchandiser.ID),
			},
		}
		handlerInstance.GetMerchandiser(ct)
		p, _ := io.ReadAll(w.Body)
		assert.Equal(t, http.StatusOK, w.Code)
		var response model.MerchandiserDB
		err = json.Unmarshal(p, &response)
		require.NoError(t, err)
		assert.Equal(t, merchandiser, response)
	})
}
