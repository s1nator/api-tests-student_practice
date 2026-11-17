package lab

import (
	ads "api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/response"
	"strings"
	"testing"
)

func TestCreateAdvertisementValidation(t *testing.T) {
	token := login(t, false)
	cases := []struct {
		name, key, value string
		omit             bool
	}{
		{"missing_title", "title", "", true}, {"missing_description", "description", "", true},
		{"empty_title", "title", "", false}, {"long_title", "title", strings.Repeat("a", 256), false},
		{"negative_price", "price", "-1", false}, {"invalid_price", "price", "abc", false},
		{"zero_quantity", "quantity", "0", false}, {"negative_quantity", "quantity", "-1", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := fields()
			if tc.omit {
				delete(f, tc.key)
			} else {
				f[tc.key] = tc.value
			}
			t.Log("Шаг: отправка невалидной формы")
			response.Error(t, ads.Create(t, token, f, photos(t), 400))
			t.Log("Результат: получена ошибка валидации")
		})
	}
	t.Run("missing_photos", func(t *testing.T) { response.Error(t, ads.Create(t, token, fields(), nil, 400)) })
	t.Run("too_many_photos", func(t *testing.T) {
		p := photos(t)[0]
		response.Error(t, ads.Create(t, token, fields(), [][]byte{p, p, p, p}, 400))
	})
}
