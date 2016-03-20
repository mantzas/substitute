package handles

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/mantzas/substitute/routes"
)

func TestAnyHandleNoContentType(t *testing.T) {

	request, _ := http.NewRequest(http.MethodGet, "/tests", nil)

	response := httptest.NewRecorder()

	AnyHandle(response, request, nil)

	if response.Code != http.StatusInternalServerError && response.Body.String() != "Content type not supported!" {
		t.Errorf("Response code should have been internal error nut was %d", response.Code)
		t.FailNow()
	}
}

func TestAnyHandleRouteNotMatched(t *testing.T) {

	request, _ := http.NewRequest(http.MethodGet, "/tests", nil)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	AnyHandle(response, request, nil)

	if response.Code != http.StatusInternalServerError && response.Body.String() != "Route not matched!" {
		t.Errorf("Response code should have been internal error nut was %d", response.Code)
		t.FailNow()
	}
}

func TestAnyHandleRouteMatched(t *testing.T) {

    routes.Register.Register(http.MethodGet, routes.JSON, "/tests", "hello world!", http.StatusCreated)

	request, _ := http.NewRequest(http.MethodGet, "/tests", nil)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	AnyHandle(response, request, nil)

	if response.Code != http.StatusCreated && response.Body.String() != "hello world!" {
		t.Errorf("Response code should have been internal error nut was %d", response.Code)
		t.FailNow()
	}
}
