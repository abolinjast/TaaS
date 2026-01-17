package store

import (
	"context"

	"github.com/abolinjast/taas/internal/models"
	"github.com/google/uuid"
)

type SessionRepository interface {
	Create(ctx context.Context, session *models.Session) error
	GetActive(ctx context.Context, userID uuid.UUID) (*models.Session, error)
	Update(ctx context.Context, session *models.Session) error
}
