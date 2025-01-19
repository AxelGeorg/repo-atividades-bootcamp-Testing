package repository

import (
	"app/internal"
	"errors"
)

type MockRepository struct {
	Products map[int]internal.Product
}

func (m *MockRepository) SearchProducts(query internal.ProductQuery) (map[int]internal.Product, error) {
	result := make(map[int]internal.Product)

	if len(m.Products) == 0 {
		return nil, errors.New("no products available")
	}

	for id, product := range m.Products {
		if query.Id == 0 || query.Id == product.Id {
			result[id] = product
		}
	}

	return result, nil
}
