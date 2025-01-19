package handler

import (
	"app/internal"
	"app/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProducts(t *testing.T) {
	mockRepo := &repository.MockRepository{
		Products: map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1", Price: 10.0, SellerId: 1}},
			2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "Product 2", Price: 20.0, SellerId: 2}},
		},
	}

	handler := NewProductsDefault(mockRepo)

	t.Run("Get existing product", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products?id=1", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.Get().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := `{"data":{"1":{"id":1,"description":"Product 1","price":10,"seller_id":1}},"message":"success"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})

	t.Run("Get non-existing product", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products?id=999", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.Get().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code for non-existing product: got %v want %v", status, http.StatusOK)
		}

		expectedEmpty := `{"data":{},"message":"success"}`
		if rr.Body.String() != expectedEmpty {
			t.Errorf("handler returned unexpected body for non-existing product: got %v want %v", rr.Body.String(), expectedEmpty)
		}
	})

	t.Run("Get product with invalid ID", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products?id=invalid", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.Get().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code for invalid ID: got %v want %v", status, http.StatusBadRequest)
		}

		expectedError := `{"status":"Bad Request","message":"invalid id"}`
		if rr.Body.String() != expectedError {
			t.Errorf("handler returned unexpected body for invalid ID: got %v want %v", rr.Body.String(), expectedError)
		}
	})

	t.Run("Get products with repository error", func(t *testing.T) {
		mockRepo2 := &repository.MockRepository{}
		handler = NewProductsDefault(mockRepo2)

		req, err := http.NewRequest(http.MethodGet, "/products?id=1", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.Get().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := `{"status":"Internal Server Error","message":"internal error"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})
}

func TestGetProductsWithStub(t *testing.T) {
	stubRepo := &repository.StubRepository{}
	handler := NewProductsDefault(stubRepo)

	t.Run("Get existing product", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products?id=1", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.Get().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := `{"data":{"1":{"id":1,"description":"Product 1","price":10,"seller_id":1}},"message":"success"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})

	t.Run("Get non-existing product", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products?id=999", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.Get().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code for non-existing product: got %v want %v", status, http.StatusOK)
		}

		expectedEmpty := `{"data":{},"message":"success"}`
		if rr.Body.String() != expectedEmpty {
			t.Errorf("handler returned unexpected body for non-existing product: got %v want %v", rr.Body.String(), expectedEmpty)
		}
	})

	t.Run("Get product with invalid ID", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products?id=invalid", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.Get().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code for invalid ID: got %v want %v", status, http.StatusBadRequest)
		}

		expectedError := `{"status":"Bad Request","message":"invalid id"}`
		if rr.Body.String() != expectedError {
			t.Errorf("handler returned unexpected body for invalid ID: got %v want %v", rr.Body.String(), expectedError)
		}
	})
}
