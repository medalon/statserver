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
	id := c.Param("id")

	act := c.QueryParam("act")
	pos := c.QueryParam("pos")

	return c.String(http.StatusOK, "id: "+id+", act: "+act+", pos: "+pos)
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
	var p []*model.Preroll

	for _, k := range stmt {
		kl := &model.Preroll{ID: k.ID, PrerollID: k.PrerollID, Name: k.Name, Date: k.Date, ShowKg: k.ShowKg, ShowWr: k.ShowWr, ClickKg: k.ClickKg, ClickWr: k.ClickWr}
		fmt.Println(kl)
		p = append(p, kl)
	}
	fmt.Println(p)
	return c.JSON(http.StatusOK, p)
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
