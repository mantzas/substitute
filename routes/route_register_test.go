package routes

import (
	"net/http"
	"testing"
)

func TestCanRegister(t *testing.T) {

	Register.Clear()

	err := Register.Register(http.MethodPost, JSON, "", "test respnse", http.StatusOK)

	if err != nil {
		t.Error("error should not have occured")
		t.FailNow()
	}

	if len(Register.routes) != 1 {
		t.Errorf("route should have been 1 but where %d", len(Register.routes))
		t.FailNow()
	}
}

func TestCanClear(t *testing.T) {

	Register.Clear()
	Register.Register(http.MethodPost, JSON, "", "test respnse", http.StatusOK)

	if len(Register.routes) != 1 {
		t.Errorf("route should have been one but where %d", len(Register.routes))
		t.FailNow()
	}

	Register.Clear()

	if len(Register.routes) != 0 {
		t.Errorf("route should have been zero but where %d", len(Register.routes))
		t.FailNow()
	}
}

func TestCanMatch(t *testing.T) {

	Register.Clear()
	Register.Register(http.MethodGet, JSON, "/users", "users response", http.StatusOK)
	Register.Register(http.MethodGet, JSON, `/users/\d`, "specific user response", http.StatusOK)
	Register.Register(http.MethodPost, JSON, "/users", "test response", http.StatusCreated)

	if len(Register.routes) != 3 {
		t.Errorf("route should have been 3 but where %d", len(Register.routes))
		t.FailNow()
	}

	matched, content, responseStatus := Register.Match(http.MethodGet, JSON, "/users/1")

	if !matched && content != "specific user response" && responseStatus != http.StatusOK {
		t.Error("Could not match route")
		t.FailNow()
	}
}
