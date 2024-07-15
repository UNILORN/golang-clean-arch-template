package shop

import (
	"context"

	shopDomain "gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/domain/shop"
)

type SaveShopUseCase struct {
	shopRepo shopDomain.ShopRepository
}

func NewSaveShopUseCase(
	shopRepo shopDomain.ShopRepository,
) *SaveShopUseCase {
	return &SaveShopUseCase{
		shopRepo: shopRepo,
	}
}

type SaveShopUseCaseInputDto struct {
	Name        string
	Description string
	Latitude    float64
	Longitude   float64
}

type SaveShopUseCaseOutputDto struct {
	ID          string
	Name        string
	Description string
	Latitude    float64
	Longitude   float64
}

func (uc *SaveShopUseCase) Run(
	ctx context.Context,
	input SaveShopUseCaseInputDto,
) (*SaveShopUseCaseOutputDto, error) {
	p, err := shopDomain.NewShop(input.Name, input.Description, input.Latitude, input.Longitude)
	if err != nil {
		return nil, err
	}
	err = uc.shopRepo.Save(ctx, p)
	if err != nil {
		return nil, err
	}
	return &SaveShopUseCaseOutputDto{
		ID:          p.GetID(),
		Name:        p.GetName(),
		Description: p.GetDescription(),
		Latitude:    p.GetLatitude(),
		Longitude:   p.GetLongitude(),
	}, nil
}
