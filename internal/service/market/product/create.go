package product

import "github.com/Ilya837/GoTgMod/internal/model/market"

func (service Service) Create(product market.Product) (uint64, error) {
	allProducts = append(allProducts, product)
	return uint64(len(allProducts)), nil
}
