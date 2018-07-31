package model

// Preroll ...
type Preroll struct {
	ID        int64  `db:"id"`
	PrerollID int64  `db:"preroll_id"`
	Name      string `db:"name"`
	Date      string `db:"date"`
	ShowKg    int64  `db:"show_kg"`
	ShowWr    int64  `db:"show_wr"`
	ClickKg   int64  `db:"click_kg"`
	ClickWr   int64  `db:"click_wr"`
}
