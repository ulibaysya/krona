package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
	"github.com/ulibaysya/krona/internal/storage/types"
)

func (db postgres) GetAllBanners() ([]types.BannerRow, error) {
	ctx := context.Background()
	const sql = "SELECT id, alias, img, redirect_url, addition_date FROM banners;"

	banners, err := queryAndCollect[types.BannerRow](ctx, db, sql, scanBanner)
	if err != nil {
		return nil, wrapError("getting all banners", err)
	}

	return banners, nil
}

func scanBanner(row pgx.CollectableRow) (types.BannerRow, error) {
	var banner types.BannerRow
	var redirectURL zeronull.Text
	if err := row.Scan(&banner.ID, &banner.Alias, &banner.Img, &redirectURL, &banner.AdditionDate); err != nil {
		return types.BannerRow{}, err
	}

	banner.RedirectURL = string(redirectURL)

	return banner, nil
}
