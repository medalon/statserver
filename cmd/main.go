package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/medalon/statserver/config"
	"github.com/medalon/statserver/stats"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Read config from system environment
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal(err, "could not get env conf parms")
	}

	s, err := stats.NewServerDB(c)
	if err != nil {
		fmt.Println(err)
	}

	e.GET("/banner/:id", s.StatBanner)
	e.GET("/preroll/:id", s.StatPreroll)

	e.GET("/getpreroll/:id", s.GetPrerollStat)
	e.GET("/getbanner/:id", s.GetBannerStat)
	e.Logger.Fatal(e.Start(":1323"))

}
