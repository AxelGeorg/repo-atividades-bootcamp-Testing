package repository

import (
	"app/internal"
	"testing"
)

func TestSearchProductsStub(t *testing.T) {
	stubRepo := &StubRepository{}

	t.Run("Search existing product by ID", func(t *testing.T) {
		query := internal.ProductQuery{Id: 1}
		results, err := stubRepo.SearchProducts(query)

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(results) != 1 {
			t.Errorf("expected 1 product, got %d", len(results))
		}
		if product, exists := results[1]; !exists || product.Id != 1 {
			t.Errorf("expected product ID 1, got %v", product)
		}
	})

	t.Run("Search non-existing product by ID", func(t *testing.T) {
		query := internal.ProductQuery{Id: 999}
		results, err := stubRepo.SearchProducts(query)

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(results) != 0 {
			t.Errorf("expected 0 products, got %d", len(results))
		}
	})

	t.Run("Search all products", func(t *testing.T) {
		query := internal.ProductQuery{}
		results, err := stubRepo.SearchProducts(query)

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(results) != 3 {
			t.Errorf("expected 3 products, got %d", len(results))
		}
	})
}
