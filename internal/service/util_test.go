package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var claimExample = jwt.MapClaims{
	"user_id":  "12",
	"username": "test@mail.com",
	"role":     "customer",
	"exp":      100,
	"iss":      "test",
}

func TestValidateJwtToken(t *testing.T) {
	claim := claimExample
	claimExample["exp"] = time.Now().Add(time.Second * time.Duration(1800)).Unix()
	secret := "sadsadewcfev"

	t.Run("Success", func(t *testing.T) {
		accessToken, err := CreateToken(claim, secret)
		require.NoError(t, err)

		token, err := GetJwtToken(accessToken, secret)
		require.NoError(t, err)

		claim, ok := token.Claims.(jwt.MapClaims)
		require.True(t, ok)
		resultClaim, err := ValidateJwtToken(accessToken, secret)
		require.NoError(t, err)

		assert.Equal(t, resultClaim, claim)
	})

	t.Run("Wrong secret", func(t *testing.T) {
		accessToken, err := CreateToken(claim, secret)
		require.NoError(t, err)

		resultClaim, err := ValidateJwtToken(accessToken, "wrong_secret")
		require.Error(t, err)
		assert.Nil(t, resultClaim)
	})
}

func TestValidateAuthHeader(t *testing.T) {
	claim := claimExample
	claimExample["exp"] = time.Now().Add(time.Second * time.Duration(1800)).Unix()
	secret := "sadsadewcfev"

	t.Run("Right header", func(t *testing.T) {
		accessToken, err := CreateToken(claim, secret)
		require.NoError(t, err)

		header := "Bearer " + accessToken
		claims, err := ValidateAuthHeader(header, secret)
		require.NoError(t, err)
		assert.NotNil(t, claims)
	})

	t.Run("Header without value", func(t *testing.T) {
		claims, err := ValidateAuthHeader("Bearer", secret)
		require.Error(t, err)
		assert.Equal(t, "wrong header", err.Error())
		assert.Nil(t, claims)
	})
}
