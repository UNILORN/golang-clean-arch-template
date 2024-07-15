package route

import (
	ginpkg "github.com/gin-gonic/gin"

	shopApp "gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/application/shop"
	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/infrastructure/mysql/query_service"
	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/infrastructure/mysql/repository"
	health_handler "gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/presentation/health_handler"
	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/presentation/settings"
	shopPre "gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/presentation/shop"
)

func InitRoute(api *ginpkg.Engine) {
	api.Use(settings.ErrorHandler())
	v1 := api.Group("/v1")
	v1.GET("/health", health_handler.HealthCheck)

	{
		shopRoute(v1)
	}
}

func shopRoute(r *ginpkg.RouterGroup) {
	shopRepository := repository.NewShopRepository()
	fetchQueryService := query_service.NewShopQueryService()
	h := shopPre.NewHandler(
		shopApp.NewSaveShopUseCase(shopRepository),
		shopApp.NewFetchShopUseCase(fetchQueryService),
	)
	group := r.Group("/shop")
	group.GET("", h.GetShops)
	group.POST("", h.PostShops)
}
