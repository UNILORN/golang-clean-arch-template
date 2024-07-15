package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/presentation/settings"

	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/application/shop"
)

type handler struct {
	saveShopUseCase  *shop.SaveShopUseCase
	fetchShopUseCase *shop.FetchShopUseCase
}

func NewHandler(
	saveShopUseCase *shop.SaveShopUseCase,
	fetchShopUseCase *shop.FetchShopUseCase,
) handler {
	return handler{
		saveShopUseCase:  saveShopUseCase,
		fetchShopUseCase: fetchShopUseCase,
	}
}

// PostShops godoc
// @Summary 商品を保存する
// @Tags shops
// @Accept json
// @Produce json
// @Param request body PostShopsParams true "登録商品"
// @Success 201 {object} postShopResponse
// @Router /v1/shops [post]
func (h handler) PostShops(ctx *gin.Context) {
	var params PostShopParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		settings.ReturnBadRequest(ctx, err)
		return
	}
	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		settings.ReturnStatusBadRequest(ctx, err)
		return
	}

	input := shop.SaveShopUseCaseInputDto{
		Name:        params.Name,
		Description: params.Description,
		Latitude:    params.Latitude,
		Longitude:   params.Longitude,
	}

	dto, err := h.saveShopUseCase.Run(ctx, input)
	if err != nil {
		settings.ReturnError(ctx, err)
		return
	}
	response := postShopResponse{
		shopResponseModel{
			Id:          dto.ID,
			Name:        dto.Name,
			Description: dto.Description,
			Latitude:    dto.Latitude,
			Longitude:   dto.Longitude,
		},
	}
	settings.ReturnStatusCreated(ctx, response)
}

// GetShops godoc
// @Summary 商品一覧を取得する
// @Tags shops
// @Accept json
// @Produce json
// @Success 200 {object} getShopsResponse
// @Router /v1/shops [get]
func (h handler) GetShops(ctx *gin.Context) {
	dtos, err := h.fetchShopUseCase.Run(ctx)
	if err != nil {
		settings.ReturnError(ctx, err)
	}

	var shops []shopResponseModel
	for _, dto := range dtos {
		shopResponseModel := shopResponseModel{
			Id:          dto.ID,
			Name:        dto.Name,
			Description: dto.Description,
			Latitude:    dto.Latitude,
			Longitude:   dto.Longitude,
		}
		shops = append(shops, shopResponseModel)
	}

	getShopsResponse := getShopsResponse{
		Shop: shops,
	}

	settings.ReturnStatusOK(ctx, getShopsResponse)
}
