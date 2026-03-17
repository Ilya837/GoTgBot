package subdomain

import "github.com/Ilya837/GoTgMod/internal/model/domain"

func (service Service) List(cursor uint64, limit uint64) ([]domain.Subdomain, error) {
	res := allSubdomains[cursor:min(cursor+limit, uint64(len(allSubdomains)))]
	return res, nil
}
