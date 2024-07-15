package shop

import (
	"context"
)

type FetchShopUseCase struct {
	shopQueryService ShopQueryService
}

func NewFetchShopUseCase(
	shopQueryService ShopQueryService,
) *FetchShopUseCase {
	return &FetchShopUseCase{
		shopQueryService: shopQueryService,
	}
}

type FetchShopUseCaseDto struct {
	ID          string
	Name        string
	Description string
	Latitude    float64
	Longitude   float64
}

func (uc *FetchShopUseCase) Run(ctx context.Context) ([]*FetchShopUseCaseDto, error) {
	qsDtos, err := uc.shopQueryService.FetchShopList(ctx)
	if err != nil {
		return nil, err
	}
	var ucDtos []*FetchShopUseCaseDto

	for _, qsDto := range qsDtos {
		ucDtos = append(ucDtos, &FetchShopUseCaseDto{
			ID:          qsDto.ID,
			Name:        qsDto.Name,
			Description: qsDto.Description,
			Latitude:    qsDto.Latitude,
			Longitude:   qsDto.Longitude,
		})
	}
	return ucDtos, err
}
