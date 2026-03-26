package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("returns API key when header is correct", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey my-secret-api-key")
		apiKey, err := GetAPIKey(headers)
		assert.NoError(t, err)
		assert.Equal(t, "my-secret-api-key", apiKey)
	})

	t.Run("returns error when header is missing", func(t *testing.T) {
		headers := http.Header{}
		_, err := GetAPIKey(headers)
		assert.ErrorIs(t, err, ErrNoAuthHeaderIncluded)
	})

	t.Run("returns error when header is malformed", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "Bearer my-secret-api-key")
		_, err := GetAPIKey(headers)
		assert.Error(t, err)
	})
}