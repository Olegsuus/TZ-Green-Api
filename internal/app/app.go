package app

import (
	"html/template"
	"log"
	"net/http"

	"TZ-GREEN-API_/internal/config"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	Config *config.Config
	Echo   *echo.Echo
}

func (a *App) Start() error {
	addr := fmt.Sprintf(":%d", a.Config.Server.Port)
	log.Printf("Starting server on %s", addr)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	t := &Template{
		templates: template.Must(template.ParseFiles("forms/index.html")),
	}
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.POST("/getSettings", a.getSettings)
	e.POST("/getStateInstance", a.getStateInstance)
	e.POST("/sendMessage", a.sendMessage)
	e.POST("/sendFileByUrl", a.sendFileByUrl)

	return e.Start(addr)
}
