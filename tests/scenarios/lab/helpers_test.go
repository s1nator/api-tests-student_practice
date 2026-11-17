package lab

import (
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/utils"
	"bytes"
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

const invalidToken = "incorrect_token"
const missingID = "00000000-0000-4000-8000-000000000000"

func TestMain(m *testing.M) { utils.LoadEnv(); os.Exit(m.Run()) }

func login(t *testing.T, second bool) string {
	t.Helper()
	prefix := "TEST_"
	if second {
		prefix = "TEST_SECOND_"
	}
	t.Log("Подготовка: авторизация тестового пользователя")
	return auth.Login(t, os.Getenv(prefix+"LOGIN"), os.Getenv(prefix+"PASSWORD")).Token
}

func fields() map[string]string {
	return map[string]string{"title": "Lab-" + utils.RandomString(16), "description": "Тестовое объявление лабораторной", "price": "125.50", "quantity": "2"}
}

func photos(t *testing.T) [][]byte {
	t.Helper()
	var b bytes.Buffer
	require.NoError(t, png.Encode(&b, image.NewRGBA(image.Rect(0, 0, 8, 8))))
	return [][]byte{b.Bytes()}
}

func create(t *testing.T, token string) (string, map[string]string) {
	t.Helper()
	f := fields()
	t.Log("Подготовка: создание уникального объявления")
	body := advertisements.Create(t, token, f, photos(t), 201)
	id := gjson.Get(body, "id").String()
	require.NotEmpty(t, id)
	t.Cleanup(func() {
		t.Log("Очистка: удаление объявления")
		advertisements.Delete(t, token, id, 204)
	})
	return id, f
}
