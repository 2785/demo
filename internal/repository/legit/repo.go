package legit

import (
	"context"

	"github.com/2785/demo/internal/models"
)

type LegitRepo struct{}

func (lr *LegitRepo) GetById(ctx context.Context, id string) (*models.Beer, error) {
	panic("this should be a legit not a mock")
}
