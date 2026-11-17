package favorites

import (
	runner "api-tests-template/internal/helpers/api-runner"
	"net/http"
	"net/url"
	"testing"
)

func Add(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return runner.GetRunner().Auth(token).Create().Post("/favorites").JSON(map[string]string{"advertisement_id": id}).Expect(t).End().Response
}

func Remove(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return runner.GetRunner().Auth(token).Create().Delete("/favorites/" + url.PathEscape(id)).Expect(t).End().Response
}
