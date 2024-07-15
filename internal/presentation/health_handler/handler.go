package health_handler

import (
	"github.com/gin-gonic/gin"

	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/presentation/settings"
)

// HealthCheck godoc
// @Summary ヘルスチェック
// @Description ヘルスチェック
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /v1/health [get]
func HealthCheck(ctx *gin.Context) {
	res := HealthResponse{
		Status: "ok",
	}
	settings.ReturnStatusOK(ctx, res)
}
