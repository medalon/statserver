package db

import "github.com/medalon/statserver/model"

// DB ...
type DB interface {
	CreatePreroll(s model.Preroll) (int64, error)
	SelectPreroll(p model.Preroll) (model.Preroll, error)
	ListPrerolls() ([]model.Preroll, error)
	UpdatePreroll(s model.Preroll) error
	DeletePreroll(id int64) error
	SelectPrerollByDate(id int64, start, end string) ([]model.Preroll, error)

	CreateBanner(s model.Banner) (int64, error)
	SelectBanner(p model.Banner) (model.Banner, error)
	ListBanners() ([]model.Banner, error)
	UpdateBanner(s model.Banner) error
	DeleteBanner(id int64) error
	SelectBannerByDate(id int64, start, end string) ([]model.Banner, error)
}
