package favorsrepo

import (
	"context"

	"github.com/nanagoboiler/models"
)

type FavorsRepository interface {
	Create(ctx context.Context, favor *models.FavorRequest) error
	CreateRelationship(ctx context.Context, req *models.RelationShipRequest) error
}
