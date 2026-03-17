package subdomain

import (
	"log"

	"github.com/Ilya837/GoTgMod/internal/model/domain"
)

func (service Service) Describe(subdomainID uint64) (*domain.Subdomain, error) {

	if int(subdomainID) > len(allSubdomains)-1 {
		log.Printf("product not exist: %d", subdomainID)
		return nil, nil
	}

	log.Printf("return product with id %d", subdomainID)
	return &allSubdomains[subdomainID], nil
}
