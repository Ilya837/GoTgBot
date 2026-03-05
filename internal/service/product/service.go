package product

import (
	"errors"
	"log"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (service Service) List() []Product {
	return allProducts
}

func (service Service) Get(id int) (*Product, error) {

	if id < 0 {
		log.Printf("wrong id: %d", id)
		return nil, errors.New("wrong input")
	}

	if id > len(allProducts)-1 {
		log.Printf("product not exist: %d", id)
		return nil, nil
	}

	log.Printf("return product: %d", id)
	return &allProducts[id], nil
}
