package model

// Banner ...
type Banner struct {
	ID       int64  `db:"id" json:"id" xml:"id"`
	BannerID int64  `db:"banner_id" json:"banner_id" xml:"banner_id"`
	Name     string `db:"name" json:"name" xml:"name"`
	Date     string `db:"date" json:"date" xml:"date"`
	ShowKg   int64  `db:"show_kg" json:"show_kg" xml:"show_kg"`
	ShowWr   int64  `db:"show_wr" json:"show_wr" xml:"show_wr"`
	ClickKg  int64  `db:"click_kg" json:"click_kg" xml:"click_kg"`
	ClickWr  int64  `db:"click_wr" json:"click_wr" xml:"click_wr"`
	Btype    string `db:"btype" json:"btype" xml:"btype"`
}
