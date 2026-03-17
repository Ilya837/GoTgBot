package subdomain

import "github.com/Ilya837/GoTgMod/internal/model/domain"

func (service Service) Create(subdomain domain.Subdomain) (uint64, error) {
	allSubdomains = append(allSubdomains, subdomain)
	return uint64(len(allSubdomains)), nil
}
