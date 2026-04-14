package product

import (
	"log"

	"github.com/Ilya837/GoTgMod/internal/model/market"
)

func (service Service) Describe(productID uint64) (*market.Product, error) {

	if int(productID) > len(allProducts)-1 {
		log.Printf("product not exist: %d", productID)
		return nil, nil
	}

	log.Printf("return product with id %d", productID)
	return &allProducts[productID], nil
}
