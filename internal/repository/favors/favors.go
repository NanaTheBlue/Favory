package favorsrepo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nanagoboiler/models"
)

type favorsRepo struct {
	pool *pgxpool.Pool
}

func NewFavorsRepository(pool *pgxpool.Pool) FavorsRepository {
	return &favorsRepo{pool: pool}
}

func (r *favorsRepo) Create(ctx context.Context, favor *models.FavorRequest) error {
	_, err := r.pool.Exec(ctx, "INSERT into favor (creator_id, recipient_id,favor_text) VALUES ($1, $2, $3);", favor.Creator_id, favor.Recipient_id, favor.Favor_text)
	if err != nil {
		return err
	}

	return nil
}

func (r *favorsRepo) CreateRelationship(ctx context.Context, req *models.RelationShipRequest) error {
	_, err := r.pool.Exec(ctx, "INSERT into relationship (requester_id, addressee_id,relationship_status) VALUES ($1, $2);", req.Inviter, req.Invitee, "pending")
	if err != nil {
		return err
	}

	return nil
}
