package product

import "github.com/Ilya837/GoTgMod/internal/model/market"

func (service Service) List(cursor uint64, limit uint64) ([]market.Product, error) {
	res := allProducts[cursor:min(cursor+limit, uint64(len(allProducts)))]
	return res, nil
}
