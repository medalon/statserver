package stats

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cactus/gostrftime"
	"github.com/labstack/echo"
	"github.com/medalon/statserver/config"
	"github.com/medalon/statserver/db"
	"github.com/medalon/statserver/model"
)

// ServerDB ...
type ServerDB struct {
	db db.DB
}

// NewServerDB ...
func NewServerDB(c *config.StatsConfig) (*ServerDB, error) {
	s := &ServerDB{}
	conn, err := db.NewMySQL(c)
	if err != nil {
		return nil, err
	}
	s.db = conn

	return s, nil
}

// StatBanner ...
func (s *ServerDB) StatBanner(c echo.Context) error {
	var geo int
	id := c.Param("id")

	act := c.QueryParam("act")
	mesto := c.QueryParam("mesto")
	name := c.QueryParam("name")
	btype := c.QueryParam("type")
	total := c.QueryParam("total")

	if mesto == "kg" {
		geo = 1
	} else {
		geo = 0
	}

	_ = s.WriteBannerToDb(id, name, act, btype, total, geo)

	return c.String(http.StatusOK, "id: "+id+", act: "+act+", name: "+name)
}

// StatPreroll ...
func (s *ServerDB) StatPreroll(c echo.Context) error {
	var geo int
	id := c.Param("id")

	act := c.QueryParam("act")
	mesto := c.QueryParam("mesto")
	name := c.QueryParam("name")

	if mesto == "kg" {
		geo = 1
	} else {
		geo = 0
	}

	_ = s.WritePrerollToDb(id, name, act, geo)

	return c.String(http.StatusOK, "id: "+id+", act: "+act+", name: "+name)
}

// GetPrerollStat ...
func (s *ServerDB) GetPrerollStat(c echo.Context) error {
	id := c.Param("id")
	rid, _ := strconv.ParseInt(id, 10, 64)
	stdate := c.QueryParam("start")
	endate := c.QueryParam("end")

	stmt, err := s.db.SelectPrerollByDate(rid, stdate, endate)

	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, stmt)
}

// GetBannerStat ...
func (s *ServerDB) GetBannerStat(c echo.Context) error {
	id := c.Param("id")
	rid, _ := strconv.ParseInt(id, 10, 64)
	stdate := c.QueryParam("start")
	endate := c.QueryParam("end")
	btype := c.QueryParam("type")
	stmt, err := s.db.SelectBannerByDate(rid, btype, stdate, endate)

	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, stmt)
}

// WritePrerollToDb ...
func (s *ServerDB) WritePrerollToDb(preid, name, act string, geo int) error {
	t1 := time.Now()
	dtime := gostrftime.Format("%Y-%m-%d", t1)

	var u model.Preroll
	u.PrerollID, _ = strconv.ParseInt(preid, 10, 64)
	u.Date = dtime

	stmt, err := s.db.SelectPreroll(u)
	switch {
	case err == sql.ErrNoRows:
		if geo == 1 {
			if act == "click" {
				u.ShowKg = 0
				u.ClickKg = 1
			} else {
				u.ShowKg = 1
				u.ClickKg = 0
			}
			u.ShowWr = 0
			u.ClickWr = 0
		} else {
			if act == "click" {
				u.ShowWr = 0
				u.ClickWr = 1
			} else {
				u.ShowWr = 1
				u.ClickWr = 0
			}
			u.ShowKg = 0
			u.ClickKg = 0
		}
		u.Name = name
		_, err := s.db.CreatePreroll(u)
		if err != nil {
			fmt.Println(err)
		}
	case err != nil:
		return err

	case stmt.ID > 0:
		if geo == 1 {
			if act == "click" {
				u.ShowKg = stmt.ShowKg
				u.ClickKg = stmt.ClickKg + 1
			} else {
				u.ShowKg = stmt.ShowKg + 1
				u.ClickKg = stmt.ClickKg
			}
			u.ShowWr = stmt.ShowWr
			u.ClickWr = stmt.ClickWr
		} else {
			if act == "click" {
				u.ShowWr = stmt.ShowWr
				u.ClickWr = stmt.ClickWr + 1
			} else {
				u.ShowWr = stmt.ShowWr + 1
				u.ClickWr = stmt.ClickWr
			}
			u.ShowKg = stmt.ShowKg
			u.ClickKg = stmt.ClickKg
		}
		u.ID = stmt.ID
		u.Name = name
		err := s.db.UpdatePreroll(u)
		if err != nil {
			return err
		}
	}

	return nil
}

// WriteBannerToDb ...
func (s *ServerDB) WriteBannerToDb(banid, name, act, btype, total string, geo int) error {
	t1 := time.Now()
	dtime := gostrftime.Format("%Y-%m-%d", t1)

	var u model.Banner
	u.BannerID, _ = strconv.ParseInt(banid, 10, 64)
	u.Date = dtime
	u.Btype = btype

	var count int64 = 1
	if total != "" {
		count, _ = strconv.ParseInt(total, 10, 64)
	}

	stmt, err := s.db.SelectBanner(u)
	switch {
	case err == sql.ErrNoRows:
		if geo == 1 {
			if act == "click" {
				u.ShowKg = 0
				u.ClickKg = count
			} else {
				u.ShowKg = count
				u.ClickKg = 0
			}
			u.ShowWr = 0
			u.ClickWr = 0
		} else {
			if act == "click" {
				u.ShowWr = 0
				u.ClickWr = count
			} else {
				u.ShowWr = count
				u.ClickWr = 0
			}
			u.ShowKg = 0
			u.ClickKg = 0
		}
		u.Name = name

		_, err := s.db.CreateBanner(u)
		if err != nil {
			fmt.Println(err)
		}
	case err != nil:
		return err

	case stmt.ID > 0:
		if geo == 1 {
			if act == "click" {
				u.ShowKg = stmt.ShowKg
				u.ClickKg = stmt.ClickKg + count
			} else {
				u.ShowKg = stmt.ShowKg + count
				u.ClickKg = stmt.ClickKg
			}
			u.ShowWr = stmt.ShowWr
			u.ClickWr = stmt.ClickWr
		} else {
			if act == "click" {
				u.ShowWr = stmt.ShowWr
				u.ClickWr = stmt.ClickWr + count
			} else {
				u.ShowWr = stmt.ShowWr + count
				u.ClickWr = stmt.ClickWr
			}
			u.ShowKg = stmt.ShowKg
			u.ClickKg = stmt.ClickKg
		}
		u.ID = stmt.ID
		u.Name = name

		err := s.db.UpdateBanner(u)
		if err != nil {
			return err
		}
	}

	return nil
}
