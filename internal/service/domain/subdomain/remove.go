package subdomain

import (
	"errors"
	"slices"
)

func (service Service) Remove(subdomainID uint64) error {

	if subdomainID > uint64(len(allSubdomains)-1) {
		return errors.New("wrong id")
	}

	allSubdomains = slices.Delete(allSubdomains, int(subdomainID), int(subdomainID)+1)

	return nil
}
