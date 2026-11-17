package lab

import (
	ads "api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/response"
	"testing"
)

func TestCreateAdvertisementUnauthorized(t *testing.T) {
	for name, token := range map[string]string{"missing_token": "", "invalid_token": invalidToken} {
		t.Run(name, func(t *testing.T) {
			t.Log("Шаг: создание без авторизации")
			response.Error(t, ads.Create(t, token, fields(), photos(t), 401))
			t.Log("Результат: доступ запрещён")
		})
	}
}
