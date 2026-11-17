# API-тесты для проекта https://testboard.avito.com/

## Полезные ссылки

1. [Swagger-спецификация](https://testboard.avito.com/swagger/index.html)
2. [Техническое задание](https://docs.google.com/document/d/1IZ9SxEcePoQT6lyEmukycp1qNok4HZtok73UQnPwKSA/edit?usp=sharing)
3. [Библиотека для написания API-тестов](https://github.com/steinfletcher/apitest)
4. [Мощная библиотека анализа JSON для Go](https://github.com/tidwall/gjson)

## Запуск тестов

1. Установите любое IDE для разработки: [Jetbrains GoLand](https://www.jetbrains.com/ru-ru/go/) (Платное или по студенческой лицензии), [VS Code](https://code.visualstudio.com/) (бесплатное).
2. Установите [Git](https://git-scm.com/install/) на компьютер
3. [Установите](https://go.dev/doc/install) Go версии 1.25.1
4. [Настройте](https://docs.github.com/ru/authentication/connecting-to-github-with-ssh) доступ к своему Github-аккаунту
   по SSH-ключу.
5. Склонируйте себе репозиторий: `git clone <путь к вашему репозиторию>`
6. Подтяните необходимые зависимости через команду в корне проекта:

```bash
   go mod tidy
```

7. В файле `.env.override` (создайте его в корне, если нет) выставьте правильные значения для переменных окружения:

```
TEST_LOGIN=your_login
TEST_PASSWORD=your_password
```

8. Для запуска всех тестов достаточно выполнить команду в корне проекта:

```bash
  go test -v ./...
```

9. Отдельный пакет с тестами можно запустить командой:

```bash
  go test -v ./tests/scenarios/myAdvertisement/
```

## Логирование

Для включения логирования метода, пути и статуса HTTP-запросов (без паролей и токенов): выставьте переменную окружения `DEBUG=true` в
`.env.override`

## Запуск линтера

1. Установить приложение [golangci-lint](https://golangci-lint.run/docs/welcome/install/local/)
2. Запуск линтера (из корня репозитория):

```bash
  golangci-lint run
```

## Доступность стенда

Перед запуском убедитесь, что открывается https://testboard.avito.com/swagger/index.html.
Если учебный стенд требует VPN, подключите его. HTTP-запросы ограничены 30 секундами.
Учётные данные храните в игнорируемом `.env.override`, не коммитьте их в `.env`.

Проверка компиляции без запросов к стенду:

```bash
go test ./... -run '^$'
```

Проверка повторяемости (при доступном стенде):

```bash
go test -v -count=3 ./...
```


## Лабораторная: объявления и избранное

В `.env.override` укажите два разных зарегистрированных аккаунта TestBoard:

```dotenv
TEST_LOGIN=owner@example.com
TEST_PASSWORD=your_password
TEST_SECOND_LOGIN=viewer@example.com
TEST_SECOND_PASSWORD=your_password
```

Тест лимита временно заполняет избранное второго аккаунта до 50 элементов
своими объявлениями и очищает их после проверки.
Доступ к стенду требует российского IP. Сценарии не регистрируют пользователей.
Каждый тест создаёт уникальные объявления и удаляет их через `t.Cleanup`;
удаление объявления также очищает связанное избранное.

```bash
go test -v -count=1 ./tests/scenarios/lab/
go test -v -count=3 ./...
golangci-lint run
```

Покрытие: создание со всеми полями и PNG, чтение, поиск, список фото;
отсутствующие title/description/photos, пустой и длинный title, отрицательная
и некорректная цена, неверное quantity, избыток фото; отсутствующий/невалидный
токен; добавление и удаление избранного, каскадное удаление, несуществующий
и некорректный id, собственное объявление и повторное добавление.

Swagger описывает структуру ошибок (`error`, `message`), но не задаёт точные
тексты: тесты проверяют статус и непустое сообщение. Ответ 500 не вызывается
искусственно. Лимит 50 проверяется отдельным сценарием `TestFavoritesLimit`.
