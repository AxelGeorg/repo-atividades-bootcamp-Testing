package repository

import (
	"app/internal"
	"testing"
)

func TestSearchProducts(t *testing.T) {
	products := map[int]internal.Product{
		1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1", Price: 10.0, SellerId: 1}},
		2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "Product 2", Price: 20.0, SellerId: 2}},
		3: {Id: 3, ProductAttributes: internal.ProductAttributes{Description: "Product 3", Price: 30.0, SellerId: 3}},
	}

	productsMap := NewProductsMap(products)

	t.Run("search existing product", func(t *testing.T) {
		query := internal.ProductQuery{Id: 1}
		results, err := productsMap.SearchProducts(query)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(results) != 1 || results[1].Id != 1 {
			t.Fatalf("expected: 1 product with ID 1, got: %d products", len(results))
		}
	})

	t.Run("search non-existing product", func(t *testing.T) {
		query := internal.ProductQuery{Id: 999}
		results, err := productsMap.SearchProducts(query)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(results) != 0 {
			t.Fatalf("expected: 0 products, got: %d products", len(results))
		}
	})

	t.Run("search all products", func(t *testing.T) {
		query := internal.ProductQuery{}
		results, err := productsMap.SearchProducts(query)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(results) != 3 {
			t.Fatalf("expected: 3 products, got: %d products", len(results))
		}
	})
}
