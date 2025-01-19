package repository

import (
	"app/internal"
	"testing"
)

func TestSearchProductsMock(t *testing.T) {
	mockRepo := &MockRepository{
		Products: map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1", Price: 10.0, SellerId: 1}},
			2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "Product 2", Price: 20.0, SellerId: 2}},
			3: {Id: 3, ProductAttributes: internal.ProductAttributes{Description: "Product 3", Price: 30.0, SellerId: 3}},
		},
	}

	t.Run("Search existing product by ID", func(t *testing.T) {
		query := internal.ProductQuery{Id: 1}
		results, err := mockRepo.SearchProducts(query)

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
		results, err := mockRepo.SearchProducts(query)

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(results) != 0 {
			t.Errorf("expected 0 products, got %d", len(results))
		}
	})

	t.Run("Search all products", func(t *testing.T) {
		query := internal.ProductQuery{}
		results, err := mockRepo.SearchProducts(query)

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(results) != 3 {
			t.Errorf("expected 3 products, got %d", len(results))
		}
	})

	t.Run("Search products with empty repository", func(t *testing.T) {
		emptyRepo := &MockRepository{Products: map[int]internal.Product{}}
		query := internal.ProductQuery{Id: 1}
		results, err := emptyRepo.SearchProducts(query)

		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		if results != nil {
			t.Errorf("expected results to be nil, got %v", results)
		}

		expectedError := "no products available"
		if err.Error() != expectedError {
			t.Errorf("expected error message %q, got %q", expectedError, err.Error())
		}
	})
}
