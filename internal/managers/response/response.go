package response

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func Read(t *testing.T, r *http.Response, status int) string {
	t.Helper()
	defer func() { require.NoError(t, r.Body.Close()) }()
	b, err := io.ReadAll(r.Body)
	require.NoError(t, err)
	require.Equal(t, status, r.StatusCode, "Неожиданный статус: %s", string(b))
	if len(b) > 0 {
		require.True(t, gjson.ValidBytes(b), "Ответ не JSON: %s", string(b))
	}
	return string(b)
}

func Error(t *testing.T, body string) {
	t.Helper()
	require.NotEmpty(t, gjson.Get(body, "error").String(), "Отсутствует код ошибки: %s", body)
	require.NotEmpty(t, gjson.Get(body, "message").String(), "Отсутствует сообщение ошибки: %s", body)
}
