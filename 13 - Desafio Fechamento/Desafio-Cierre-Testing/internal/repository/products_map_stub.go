package repository

import "app/internal"

type StubRepository struct{}

func (s *StubRepository) SearchProducts(query internal.ProductQuery) (map[int]internal.Product, error) {
	fixedProducts := map[int]internal.Product{
		1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1", Price: 10.0, SellerId: 1}},
		2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "Product 2", Price: 20.0, SellerId: 2}},
		3: {Id: 3, ProductAttributes: internal.ProductAttributes{Description: "Product 3", Price: 30.0, SellerId: 3}},
	}

	result := make(map[int]internal.Product)
	for id, product := range fixedProducts {
		if query.Id == 0 || query.Id == product.Id {
			result[id] = product
		}
	}
	return result, nil
}
