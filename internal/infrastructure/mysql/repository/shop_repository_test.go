package repository

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/oklog/ulid/v2"

	shopDomain "gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/domain/shop"
)

func Test_shopRepository_Save_And_Find(t *testing.T) {
	testStr := "test"
	shopID1 := ulid.Make().String()
	shopID2 := ulid.Make().String()
	shop1, _ := shopDomain.Reconstruct(
		shopID1,
		testStr,
		testStr,

		12.3333,
		104.1111,
	)
	shop2, _ := shopDomain.Reconstruct(
		shopID2,
		testStr,
		testStr,
		12.3333,
		104.1111,
	)
	tests := []struct {
		name    string
		shop    []*shopDomain.Shop
		wantErr bool
	}{
		{
			name:    "正常系",
			shop:    []*shopDomain.Shop{shop1, shop2},
			wantErr: false,
		},
	}
	repo := NewShopRepository()
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// shopを全て保存
			for _, p := range tt.shop {
				if err := repo.Save(ctx, p); (err != nil) != tt.wantErr {
					t.Errorf("shopRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

			// IDで検索
			p, err := repo.FindByID(ctx, tt.shop[0].GetID())
			if err != nil {
				t.Errorf("shopRepository.FindByID() error = %v", err)
			}
			if diff := cmp.Diff(
				tt.shop[0],
				p,
				cmp.AllowUnexported(shopDomain.Shop{}),
			); diff != "" {
				t.Errorf("shopRepository.FindByID() = %v, want %v. error is %s", p, tt.shop, err)
			}

			// IDsで検索
			ps, err := repo.FindByIDs(ctx, []string{shopID1, shopID2})
			if err != nil {
				t.Errorf("shopRepository.FindByIDs() error = %v", err)
			}
			if diff := cmp.Diff(
				tt.shop,
				ps,
				cmp.AllowUnexported(shopDomain.Shop{}),
			); diff != "" {
				t.Errorf("shopRepository.FindByIDs() = %v, want %v. error is %s", ps, tt.shop, err)
			}
		})
	}
}
