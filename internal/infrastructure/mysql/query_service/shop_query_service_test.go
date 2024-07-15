package query_service

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/application/shop"
)

func Test_Fetch_Product_Query_Service(t *testing.T) {
	p := []*shop.FetchShopListDto{
		{
			ID:          "01HCNYK4MQNC6G6X3F3DGXZ2J8",
			Name:        "サウナハット",
			Description: "サウナハット",
			Latitude:    35.6895,
			Longitude:   139.6917,
		},
	}
	tests := []struct {
		name string
		want []*shop.FetchShopListDto
	}{
		{
			name: "オーナ情報を含めた商品一覧が取得できる",
			want: p,
		},
	}

	queryService := NewShopQueryService()
	resetTestData(t)
	for _, tt := range tests {

		t.Run(fmt.Sprintf(": %s", tt.name), func(t *testing.T) {
			got, _ := queryService.FetchShopList(context.Background())
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FindById() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
