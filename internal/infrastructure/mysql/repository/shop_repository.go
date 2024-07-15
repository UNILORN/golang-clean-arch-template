package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	errDomain "gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/domain/error"
	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/domain/shop"
	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/infrastructure/mysql/db"
	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/infrastructure/mysql/db/dbgen"
)

type shopRepository struct {
}

func NewShopRepository() shop.ShopRepository {
	return &shopRepository{}
}

func (r *shopRepository) Save(ctx context.Context, shop *shop.Shop) error {
	query := db.GetQuery(ctx)
	if err := query.UpsertShop(ctx, dbgen.UpsertShopParams{
		ID:          shop.GetID(),
		Name:        shop.GetName(),
		Description: shop.GetDescription(),
	}); err != nil {
		return err
	}
	return nil
}

func (r *shopRepository) FindByID(ctx context.Context, id string) (*shop.Shop, error) {
	query := db.GetQuery(ctx)
	p, err := query.ShopFindById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDomain.NotFoundErr
		}
		return nil, err
	}

	lat, err := strconv.ParseFloat(p.Latitude, 64)
	if err != nil {
		return nil, err
	}

	lon, err := strconv.ParseFloat(p.Longitude, 64)
	if err != nil {
		return nil, err
	}

	pd, err := shop.Reconstruct(
		p.ID,
		p.Name,
		p.Description,
		lat,
		lon,
	)
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (r *shopRepository) FindByIDs(ctx context.Context, ids []string) ([]*shop.Shop, error) {
	query := db.GetQuery(ctx)
	ps, err := query.ShopFindByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	var products []*shop.Shop
	for _, p := range ps {

		lat, err := strconv.ParseFloat(p.Latitude, 64)
		if err != nil {
			return nil, err
		}

		lon, err := strconv.ParseFloat(p.Longitude, 64)
		if err != nil {
			return nil, err
		}

		pd, err := shop.Reconstruct(
			p.ID,
			p.Name,
			p.Description,
			lat,
			lon,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, pd)
	}
	return products, nil
}
