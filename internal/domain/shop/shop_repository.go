//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package shop

import (
	"context"
)

type ShopRepository interface {
	Save(ctx context.Context, product *Shop) error
	FindByID(ctx context.Context, id string) (*Shop, error)
	FindByIDs(ctx context.Context, ids []string) ([]*Shop, error)
}
