package lab

import (
	fav "api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/managers/response"
	"testing"
)

func TestFavoritesUnauthorized(t *testing.T) {
	for name, token := range map[string]string{"missing_token": "", "invalid_token": invalidToken} {
		t.Run(name, func(t *testing.T) {
			t.Run("list", func(t *testing.T) { response.Error(t, fav.List(t, token, "", 401)) })
			t.Run("add", func(t *testing.T) { response.Error(t, fav.Add(t, token, missingID, 401)) })
			t.Run("remove", func(t *testing.T) { response.Error(t, fav.Remove(t, token, missingID, 401)) })
		})
	}
}
