//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package shop

import "context"

type FetchShopListDto struct {
	ID          string
	Name        string
	Description string
	Latitude    float64
	Longitude   float64
}

type ShopQueryService interface {
	FetchShopList(ctx context.Context) ([]*FetchShopListDto, error)
}
