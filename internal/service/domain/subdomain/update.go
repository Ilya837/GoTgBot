package subdomain

import (
	"errors"

	"github.com/Ilya837/GoTgMod/internal/model/domain"
)

func (service Service) Update(subdomainID uint64, subdomain domain.Subdomain) error {

	if subdomainID > uint64(len(allSubdomains)-1) {
		return errors.New("wrong id")
	}

	allSubdomains[subdomainID] = subdomain

	return nil
}
