package product

import (
	"github.com/Ilya837/GoTgMod/internal/model/market"
)

var allProducts = []market.Product{

	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
}

type ProductService interface {
	Describe(productID uint64) (*market.Product, error)
	List(cursor uint64, limit uint64) ([]market.Product, error)
	Create(product market.Product) (uint64, error)
	Update(productID uint64, product market.Product) error
	Remove(productID uint64) error
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

type DummyProductService struct{}

func NewDummyProductService() *DummyProductService {
	return &DummyProductService{}
}
