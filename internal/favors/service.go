package favors

import (
	"context"

	authrepo "github.com/nanagoboiler/internal/repository/auth"
	favorsrepo "github.com/nanagoboiler/internal/repository/favors"
	"github.com/nanagoboiler/models"
)

type favorService struct {
	FavorsRepo favorsrepo.FavorsRepository
	UserRepo   authrepo.UserRepository
}

func NewFavorService(userrepo authrepo.UserRepository, favorsrepo favorsrepo.FavorsRepository) Service {
	return &favorService{UserRepo: userrepo, FavorsRepo: favorsrepo}
}

func (s *favorService) VerifyRelationship(ctx context.Context, Creator_id string, Recipient_id string) (bool, error) {

	return true, nil
}

func (s *favorService) CreateRelationship(ctx context.Context, req *models.RelationShipRequest) error {

	err := s.FavorsRepo.CreateRelationship(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (s *favorService) CreateFavor(ctx context.Context, req *models.FavorRequest) error {

	err := s.FavorsRepo.Create(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
