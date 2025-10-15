package favorsrepo

import (
	"context"

	"github.com/nanagoboiler/models"
)

type FavorsRepository interface {
	Create(ctx context.Context, favor *models.FavorRequest) error
}
