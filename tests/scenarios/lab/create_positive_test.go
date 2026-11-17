package lab

import (
	ads "api-tests-template/internal/managers/advertisements"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestCreateAdvertisementPositive(t *testing.T) {
	token := login(t, false)
	id, f := create(t, token)
	t.Log("Шаг: получение объявления и проверка всех переданных полей")
	body := ads.Get(t, id)
	for _, key := range []string{"title", "description"} {
		require.Equal(t, f[key], gjson.Get(body, key).String(), key)
	}
	require.Equal(t, 125.5, gjson.Get(body, "price").Float())
	require.Equal(t, int64(2), gjson.Get(body, "quantity").Int())
	t.Log("Шаг: поиск по полному названию")
	items := gjson.Get(ads.Search(t, f["title"]), "items").Array()
	require.Len(t, items, 1)
	require.Equal(t, id, items[0].Get("id").String())
	t.Log("Шаг: проверка загруженной фотографии")
	listed := gjson.Parse(ads.Photos(t, token, id)).Array()
	require.Len(t, listed, 1)
	require.NotEmpty(t, listed[0].Get("id").String())
	require.NotEmpty(t, listed[0].Get("url").String())
	require.Equal(t, gjson.Get(body, "photos.0.id").String(), listed[0].Get("id").String())
	t.Log("Результат: поля, поиск и фотография проверены")
}
