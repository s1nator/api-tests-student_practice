package lab

import (
	ads "api-tests-template/internal/managers/advertisements"
	fav "api-tests-template/internal/managers/favorites"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestFavoritesDeletedAdvertisement(t *testing.T) {
	owner := login(t, false)
	viewer := login(t, true)
	f := fields()
	body := ads.Create(t, owner, f, photos(t), 201)
	id := gjson.Get(body, "id").String()
	require.NotEmpty(t, id)
	deleted := false
	t.Cleanup(func() {
		if !deleted {
			ads.Delete(t, owner, id, 204)
		}
	})
	fav.Add(t, viewer, id, 201)
	require.Len(t, gjson.Get(fav.List(t, viewer, f["title"], 200), "items").Array(), 1)
	t.Log("Шаг: удаление самого объявления")
	ads.Delete(t, owner, id, 204)
	deleted = true
	require.Empty(t, gjson.Get(fav.List(t, viewer, f["title"], 200), "items").Array())
	t.Log("Результат: объявление исчезло из избранного")
}
