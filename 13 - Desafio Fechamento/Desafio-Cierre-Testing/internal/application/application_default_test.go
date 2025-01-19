package application

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplicationDefault_SetUp(t *testing.T) {
	cfg := &ConfigApplicationDefault{
		Addr: ":8080",
	}
	app := NewApplicationDefault(cfg)

	err := app.SetUp()
	if err != nil {
		t.Fatalf("expected no error during setup, got %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, "/product/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	app.rt.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status OK, got %v", status)
	}

	expected := `{"data":{},"message":"success"}`
	if rr.Body.String() != expected {
		t.Errorf("unexpected response body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestApplicationDefault_Run(t *testing.T) {
	cfg := &ConfigApplicationDefault{
		Addr: ":8080",
	}
	app := NewApplicationDefault(cfg)

	err := app.SetUp()
	if err != nil {
		t.Fatalf("expected no error during setup, got %v", err)
	}

	go func() {
		if err := app.Run(); err != nil {
			t.Fatalf("expected no error when running application, got %v", err)
		}
	}()

	req, err := http.NewRequest(http.MethodGet, "/product/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	app.rt.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status OK, got %v", status)
	}
}

func TestApplicationDefault_TearDown(t *testing.T) {
	cfg := &ConfigApplicationDefault{
		Addr: ":8080",
	}
	app := NewApplicationDefault(cfg)

	err := app.TearDown()
	if err != nil {
		t.Fatalf("expected no error during teardown, got %v", err)
	}
}
