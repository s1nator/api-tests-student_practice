package advertisementssearch

import (
	"bytes"
	"image"
	"image/png"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"api-tests-template/internal/assertions"
	advList "api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
)

func TestSearchByTitleSubstring(t *testing.T) {
	base.SetupSuite()
	t.Log("Подготовка: авторизация и создание уникального объявления")
	token := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD")).Token
	beginning := utils.RandomString(10)
	middle := "Вулкан-" + utils.RandomString(10)
	end := utils.RandomString(10)
	title := beginning + "-" + middle + "-" + end
	var photo bytes.Buffer
	require.NoError(t, png.Encode(&photo, image.NewRGBA(image.Rect(0, 0, 8, 8))))
	body := advList.Create(t, token, map[string]string{"title": title, "description": "Проверка поиска по подстроке", "price": "10", "quantity": "1"}, [][]byte{photo.Bytes()}, 201)
	id := gjson.Get(body, "id").String()
	require.NotEmpty(t, id)
	t.Cleanup(func() {
		t.Log("Очистка: удаление тестового объявления")
		advList.Delete(t, token, id, 204)
	})
	cases := []struct{ name, search string }{
		{"начало", beginning},
		{"середина", middle},
		{"конец", end},
		{"нижний регистр", strings.ToLower(middle)},
		{"верхний регистр", strings.ToUpper(middle)},
		{"смешанный регистр", utils.MixCase(middle)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("Шаг: поиск по подстроке названия")
			list := advList.GetAdvertisementsList(t, token, tc.search)
			assertions.ContainsAdvertisementByTitle(t, list.Items, title)
			t.Log("Результат: созданное объявление найдено")
		})
	}
}
