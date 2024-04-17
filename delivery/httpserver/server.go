package httpserver

import (
	"app/config"
	"app/delivery/httpserver/discounthandler"
	"app/service/discountservice"
	"app/validator/discountvalidator"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config          config.Config
	Router          *echo.Echo
	discountHandler discounthandler.Handler
}

func New(config config.Config, discountSvc discountservice.Service, discountValidator discountvalidator.Validator) Server {
	return Server{
		config:          config,
		Router:          echo.New(),
		discountHandler: discounthandler.New(discountSvc, discountValidator),
	}
}

func (s Server) Serve() *echo.Echo {
	// Middleware
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())
	s.discountHandler.SetUserRoutes(s.Router)
	address := fmt.Sprintf(":%d", s.config.HTTPServer.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
	return s.Router
}
