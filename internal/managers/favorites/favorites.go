package favorites

import (
	ads "api-tests-template/internal/client/http/advertisements"
	client "api-tests-template/internal/client/http/favorites"
	"api-tests-template/internal/managers/response"
	"testing"
)

func Add(t *testing.T, token, id string, status int) string {
	t.Helper()
	return response.Read(t, client.Add(t, token, id), status)
}

func Remove(t *testing.T, token, id string, status int) string {
	t.Helper()
	return response.Read(t, client.Remove(t, token, id), status)
}

func List(t *testing.T, token, title string, status int) string {
	t.Helper()
	return response.Read(t, ads.Get(t, token, "/favorites", map[string]string{"search": title, "page_size": "50"}), status)
}
