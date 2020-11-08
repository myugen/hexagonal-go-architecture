package types

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
)

type Datetime struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

var _ pg.BeforeUpdateHook = (*Datetime)(nil)

func (b *Datetime) BeforeUpdate(ctx context.Context) (context.Context, error) {
	b.UpdatedAt = time.Now()
	return ctx, nil
}

type SoftDelete struct {
	DeletedAt time.Time `pg:",soft_delete"`
}
