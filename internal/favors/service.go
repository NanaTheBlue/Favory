package favors

import (
	"context"

	favorsrepo "github.com/nanagoboiler/internal/repository/favors"
	"github.com/nanagoboiler/models"
)

type favorService struct {
	FavorsRepo favorsrepo.FavorsRepository
}

func NewFavorService(favorsrepo favorsrepo.FavorsRepository) Service {
	return &favorService{FavorsRepo: favorsrepo}
}

func (s *favorService) CreateFavor(ctx context.Context, req *models.FavorRequest) error {

	err := s.FavorsRepo.Create(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
