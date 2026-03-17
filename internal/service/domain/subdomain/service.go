package subdomain

import (
	"github.com/Ilya837/GoTgMod/internal/model/domain"
)

var allSubdomains = []domain.Subdomain{

	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
}

type SubdomainService interface {
	Describe(subdomainID uint64) (*domain.Subdomain, error)
	List(cursor uint64, limit uint64) ([]domain.Subdomain, error)
	Create(subdomain domain.Subdomain) (uint64, error)
	Update(subdomainID uint64, subdomain domain.Subdomain) error
	Remove(subdomainID uint64) error
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

type DummySubdomainService struct{}

func NewDummySubdomainService() *DummySubdomainService {
	return &DummySubdomainService{}
}
