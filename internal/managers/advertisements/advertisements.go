package advertisements

import (
	client "api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/managers/response"
	"net/url"
	"testing"
)

func Create(t *testing.T, token string, fields map[string]string, photos [][]byte, status int) string {
	t.Helper()
	return response.Read(t, client.Create(t, token, fields, photos), status)
}

func Get(t *testing.T, id string) string {
	t.Helper()
	return response.Read(t, client.Get(t, "", "/advertisement", map[string]string{"id": id}), 200)
}

func Search(t *testing.T, title string) string {
	t.Helper()
	return response.Read(t, client.Get(t, "", "/advertisements", map[string]string{"search": title, "limit": "50"}), 200)
}

func Photos(t *testing.T, token, id string) string {
	t.Helper()
	return response.Read(t, client.Get(t, token, "/advertisements/"+url.PathEscape(id)+"/photos", nil), 200)
}

func Delete(t *testing.T, token, id string, status int) {
	t.Helper()
	response.Read(t, client.Delete(t, token, id), status)
}
