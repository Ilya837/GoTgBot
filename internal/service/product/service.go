package product

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (service Service) List() []Product {
	return allProducts
}

func (service Service) Get(id int) (*Product, error) {
	//добавить обработку
	return &allProducts[id], nil
}
