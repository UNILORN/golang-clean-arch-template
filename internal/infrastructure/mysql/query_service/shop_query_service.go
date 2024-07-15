package query_service

import (
	"context"
	"strconv"

	shopQS "gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/application/shop"
	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/infrastructure/mysql/db"
)

type shopQueryService struct{}

func NewShopQueryService() shopQS.ShopQueryService {
	return &shopQueryService{}
}

func (q *shopQueryService) FetchShopList(ctx context.Context) ([]*shopQS.FetchShopListDto, error) {
	query := db.GetReadQuery()
	shopList, err := query.ShopFetch(ctx)
	if err != nil {
		return nil, err
	}

	var shopFetchServiceDtos []*shopQS.FetchShopListDto
	for _, shop := range shopList {
		lat, err := strconv.ParseFloat(shop.Latitude, 64)
		if err != nil {
			return nil, err
		}

		lon, err := strconv.ParseFloat(shop.Longitude, 64)
		if err != nil {
			return nil, err
		}

		shopFetchServiceDtos = append(shopFetchServiceDtos, &shopQS.FetchShopListDto{
			ID:          shop.ID,
			Name:        shop.Name,
			Description: shop.Description,
			Latitude:    lat,
			Longitude:   lon,
		})
	}
	return shopFetchServiceDtos, nil
}
