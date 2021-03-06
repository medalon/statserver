package db

import (

	// This line is must for working MySQL database
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/medalon/statserver/config"
	"github.com/medalon/statserver/model"
)

// MySQL provides api for work with mysql database
type MySQL struct {
	conn *sqlx.DB
}

// NewMySQL creates a new instance of database API
func NewMySQL(c *config.StatsConfig) (*MySQL, error) {
	conn, err := sqlx.Open("mysql", c.DatabaseURL)
	if err != nil {
		return nil, err
	}

	m := &MySQL{}
	m.conn = conn
	return m, nil
}

// CreatePreroll creates preroll entry in database
func (m *MySQL) CreatePreroll(s model.Preroll) (int64, error) {
	res, err := m.conn.Exec(
		"INSERT INTO `prerolls` (`preroll_id`, `name`, `date`, `show_kg`, `show_wr`, `click_kg`, `click_wr`) VALUES (?, ?, ?, ?, ?, ?, ?)", s.PrerollID, s.Name, s.Date, s.ShowKg, s.ShowWr, s.ClickKg, s.ClickWr,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// SelectPreroll selects preroll entry from database
func (m *MySQL) SelectPreroll(p model.Preroll) (model.Preroll, error) {
	var s model.Preroll
	err := m.conn.Get(&s, "SELECT `id`, `show_kg`, `show_wr`, `click_kg`, `click_wr` FROM `prerolls` WHERE `preroll_id`=? AND `date`=?", p.PrerollID, p.Date)
	return s, err
}

// ListPrerolls returns array of prerolls entries from database
func (m *MySQL) ListPrerolls() ([]model.Preroll, error) {
	prerolls := []model.Preroll{}
	err := m.conn.Select(&prerolls, "SELECT * FROM `prerolls`")
	return prerolls, err
}

// UpdatePreroll updates preroll entry in database
func (m *MySQL) UpdatePreroll(s model.Preroll) error {
	tx := m.conn.MustBegin()
	tx.MustExec(
		"UPDATE `prerolls` SET `preroll_id` = ?, `name` = ?, `date` = ?, `show_kg` = ?, `show_wr` = ?, `click_kg` = ?, `click_wr` = ? WHERE `id` = ?",
		s.PrerollID, s.Name, s.Date, s.ShowKg, s.ShowWr, s.ClickKg, s.ClickWr, s.ID,
	)
	err := tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

// DeletePreroll deletes preroll entry from database
func (m *MySQL) DeletePreroll(id int64) error {
	_, err := m.conn.Exec("DELETE FROM `prerolls` WHERE id=?", id)
	return err
}

// SelectPrerollByDate selects preroll entry from database
func (m *MySQL) SelectPrerollByDate(id int64, start, end string) ([]model.Preroll, error) {
	s := []model.Preroll{}
	err := m.conn.Select(&s, "SELECT * FROM `prerolls` WHERE `preroll_id`=? AND `date`>=? AND `date`<=?", id, start, end)
	return s, err
}

// *** for banners *** //

// CreateBanner creates banner entry in database
func (m *MySQL) CreateBanner(s model.Banner) (int64, error) {
	res, err := m.conn.Exec(
		"INSERT INTO `banners` (`banner_id`, `name`, `date`, `show_kg`, `show_wr`, `click_kg`, `click_wr`, `btype`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", s.BannerID, s.Name, s.Date, s.ShowKg, s.ShowWr, s.ClickKg, s.ClickWr, s.Btype,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// SelectBanner selects banner entry from database
func (m *MySQL) SelectBanner(p model.Banner) (model.Banner, error) {
	var s model.Banner
	err := m.conn.Get(&s, "SELECT `id`, `show_kg`, `show_wr`, `click_kg`, `click_wr`, `btype` FROM `banners` WHERE `banner_id`=? AND `date`=? AND `btype`=?", p.BannerID, p.Date, p.Btype)
	return s, err
}

// ListBanners returns array of banners entries from database
func (m *MySQL) ListBanners() ([]model.Banner, error) {
	banners := []model.Banner{}
	err := m.conn.Select(&banners, "SELECT * FROM `banners`")
	return banners, err
}

// UpdateBanner updates banner entry in database
func (m *MySQL) UpdateBanner(s model.Banner) error {
	tx := m.conn.MustBegin()
	tx.MustExec(
		"UPDATE `banners` SET `banner_id` = ?, `name` = ?, `date` = ?, `show_kg` = ?, `show_wr` = ?, `click_kg` = ?, `click_wr` = ?, `btype` = ? WHERE `id` = ?",
		s.BannerID, s.Name, s.Date, s.ShowKg, s.ShowWr, s.ClickKg, s.ClickWr, s.Btype, s.ID,
	)
	err := tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

// DeleteBanner deletes banner entry from database
func (m *MySQL) DeleteBanner(id int64) error {
	_, err := m.conn.Exec("DELETE FROM `banners` WHERE id=?", id)
	return err
}

// SelectBannerByDate selects banners entries from database
func (m *MySQL) SelectBannerByDate(id int64, btype, start, end string) ([]model.Banner, error) {
	s := []model.Banner{}
	err := m.conn.Select(&s, "SELECT * FROM `banners` WHERE `banner_id`=? AND `btype`=? AND `date`>=? AND `date`<=?", id, btype, start, end)
	return s, err
}
