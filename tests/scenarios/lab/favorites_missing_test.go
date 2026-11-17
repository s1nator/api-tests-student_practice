package lab

import (
	fav "api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/managers/response"
	"testing"
)

func TestFavoritesMissingAdvertisement(t *testing.T) {
	token := login(t, true)
	t.Run("add", func(t *testing.T) {
		t.Log("Шаг: добавление несуществующего id")
		response.Error(t, fav.Add(t, token, missingID, 404))
	})
	t.Run("remove", func(t *testing.T) {
		t.Log("Шаг: удаление несуществующего id")
		response.Error(t, fav.Remove(t, token, missingID, 404))
	})
}
