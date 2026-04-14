package product

import (
	"errors"

	"github.com/Ilya837/GoTgMod/internal/model/market"
)

func (service Service) Update(productID uint64, product market.Product) error {

	if productID > uint64(len(allProducts)-1) {
		return errors.New("wrong id")
	}

	allProducts[productID] = product

	return nil
}
