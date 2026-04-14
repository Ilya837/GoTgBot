package product

import (
	"errors"
	"slices"
)

func (service Service) Remove(productID uint64) error {

	if productID > uint64(len(allProducts)-1) {
		return errors.New("wrong id")
	}

	allProducts = slices.Delete(allProducts, int(productID), int(productID)+1)

	return nil
}
