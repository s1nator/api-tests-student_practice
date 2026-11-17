package lab

import (
	fav "api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/managers/response"
	"testing"
)

func TestFavoritesOwnAdvertisement(t *testing.T) {
	token := login(t, false)
	id, _ := create(t, token)
	t.Log("Шаг: добавление своего объявления")
	response.Error(t, fav.Add(t, token, id, 403))
}
func TestFavoritesDuplicate(t *testing.T) {
	owner := login(t, false)
	viewer := login(t, true)
	id, _ := create(t, owner)
	fav.Add(t, viewer, id, 201)
	t.Log("Шаг: повторное добавление")
	response.Error(t, fav.Add(t, viewer, id, 409))
}
func TestFavoritesInvalidID(t *testing.T) {
	token := login(t, true)
	t.Log("Шаг: некорректный id")
	response.Error(t, fav.Add(t, token, "not-a-uuid", 400))
}
