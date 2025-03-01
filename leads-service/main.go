package main

import (
	"leadsservice/config"
	"leadsservice/controllers"
	"leadsservice/repositories"
	"leadsservice/services"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	config.InitDB()
	// Create instances of repositories and services
	leadRepo := repositories.NewLeadsRepository()
	leadService := services.NewLeadService(leadRepo)
	leadControllerr := controllers.NewLeadsController(leadService, leadRepo)
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))
	e.POST("/api/leads", leadControllerr.Submit)
	e.GET("/api/leads", leadControllerr.Leads)
	e.GET("/api/leads/:id", leadControllerr.LeadsByID)

	e.Logger.Fatal(e.Start(":8181"))
}

// Struct untuk menangani validasi
type CustomValidator struct {
	validator *validator.Validate
}

// Implementasi metode `Validate` untuk CustomValidator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
