package favors

import (
	"context"

	"github.com/nanagoboiler/models"
)

type Service interface {
	CreateFavor(ctx context.Context, req *models.FavorRequest) error
}
