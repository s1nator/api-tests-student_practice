package lab

import (
	fav "api-tests-template/internal/managers/favorites"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestFavoritesCycle(t *testing.T) {
	owner := login(t, false)
	viewer := login(t, true)
	id, f := create(t, owner)
	t.Log("Шаг: добавление чужого объявления в избранное")
	added := fav.Add(t, viewer, id, 201)
	require.Equal(t, id, gjson.Get(added, "advertisement.id").String())

	items := gjson.Get(fav.List(t, viewer, f["title"], 200), "items").Array()
	require.Len(t, items, 1)
	require.Equal(t, id, items[0].Get("advertisement.id").String())
	t.Log("Шаг: удаление из избранного")
	fav.Remove(t, viewer, id, 204)
	require.Empty(t, gjson.Get(fav.List(t, viewer, f["title"], 200), "items").Array())
	t.Log("Результат: полный цикл избранного проверен")
}
