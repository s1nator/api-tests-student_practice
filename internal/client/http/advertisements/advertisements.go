package advertisements

import (
	runner "api-tests-template/internal/helpers/api-runner"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func Create(t *testing.T, token string, fields map[string]string, photos [][]byte) *http.Response {
	t.Helper()
	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	for k, v := range fields {
		require.NoError(t, w.WriteField(k, v))
	}
	for _, photo := range photos {
		part, err := w.CreateFormFile("photos", "test.png")
		require.NoError(t, err)
		_, err = io.Copy(part, bytes.NewReader(photo))
		require.NoError(t, err)
	}
	require.NoError(t, w.Close())
	return runner.GetRunner().Auth(token).Create().Post("/advertisement").ContentType(w.FormDataContentType()).Body(body.String()).Expect(t).End().Response
}

func Get(t *testing.T, token, path string, query map[string]string) *http.Response {
	t.Helper()
	req := runner.GetRunner().Auth(token).Create().Get(path)
	for k, v := range query {
		req.Query(k, v)
	}
	return req.Expect(t).End().Response
}

func Delete(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return runner.GetRunner().Auth(token).Create().Delete("/advertisement").Query("id", id).Expect(t).End().Response
}
