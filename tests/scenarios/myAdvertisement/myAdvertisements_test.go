package myAdvertisement

import (
	ads "api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/utils"
	"bytes"
	"image"
	"image/png"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"

	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/managers/myAdvertisements"

	base "api-tests-template/tests"
)

type TestSuite struct {
	suite.Suite
	loginData models.LoginOkResponse
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()

	base.Precondition("Авторизация пользователя с кредами из переменных окружения и получение его параметров")
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

func (s *TestSuite) TestGetMyAdvertisementsPositive() {
	t := s.T()
	t.Log("Подготовка: создание собственного объявления")
	var photo bytes.Buffer
	require.NoError(t, png.Encode(&photo, image.NewRGBA(image.Rect(0, 0, 8, 8))))
	title := "MyAds-" + utils.RandomString(16)
	body := ads.Create(t, s.loginData.Token, map[string]string{"title": title, "description": "Тестовое объявление", "price": "10", "quantity": "1"}, [][]byte{photo.Bytes()}, http.StatusCreated)
	id := gjson.Get(body, "id").String()
	require.NotEmpty(t, id)
	t.Cleanup(func() { ads.Delete(t, s.loginData.Token, id, http.StatusNoContent) })
	t.Log("Шаг: получение собственных объявлений")
	items := gjson.Get(myAdvertisements.GetMyAdvertisements(t, s.loginData.Token, http.StatusOK, 50), "items").Array()
	found := false
	for _, item := range items {
		if item.Get("id").String() == id {
			found = true
			require.Equal(t, s.loginData.User.Id, item.Get("user_id").String())
			require.Equal(t, title, item.Get("title").String())
			require.NotEmpty(t, item.Get("photos").Array())
		}
	}
	require.True(t, found, "Созданное объявление отсутствует в моих объявлениях")
	t.Log("Результат: собственное объявление найдено")
}

func (s *TestSuite) TestGetMyAdvertisementsIncorrectToken() {
	var advertisementsBody string
	s.Run("Получаем список собственных объявлений с неправильным авторизационным токеном", func() {
		advertisementsBody = myAdvertisements.GetMyAdvertisements(s.T(), "incorrect_token", http.StatusUnauthorized, 50)
	})

	s.Run("Проверяем, что показывается правильная ошибка", func() {
		require.Equal(s.T(), "unauthorized", gjson.Get(advertisementsBody, "error").String(),
			"Некорректное значение error параметра")
		require.Equal(s.T(), "Invalid or expired token", gjson.Get(advertisementsBody, "message").String(),
			"Некорректное значение message параметра")
	})
}
