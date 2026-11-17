package lab

import (
	fav "api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/managers/response"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestFavoritesLimit(t *testing.T) {
	owner := login(t, false)
	viewer := login(t, true)
	total := gjson.Get(fav.List(t, viewer, "", 200), "total").Int()
	require.LessOrEqual(t, total, int64(50), "Исходное избранное превышает лимит API")
	t.Log("Подготовка: заполнение свободных мест избранного тестовыми объявлениями")
	for i := total; i < 50; i++ {
		id, _ := create(t, owner)
		fav.Add(t, viewer, id, 201)
	}
	id, _ := create(t, owner)
	t.Log("Шаг: добавление 51-го объявления")
	response.Error(t, fav.Add(t, viewer, id, 400))
	require.Equal(t, int64(50), gjson.Get(fav.List(t, viewer, "", 200), "total").Int())
	t.Log("Результат: лимит избранного соблюдается")
}
