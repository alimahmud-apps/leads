package controllers

import (
	"fmt"
	"leadsservice/models"
	"leadsservice/repositories"
	"leadsservice/services"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type LeadsController struct {
	leadService services.LeadService
	leadRepo    repositories.LeadRepository
}

func NewLeadsController(leadService services.LeadService, leadRepo repositories.LeadRepository) *LeadsController {
	return &LeadsController{
		leadService: leadService,
		leadRepo:    leadRepo,
	}
}

// customErrorMessage custom func to mapping error validate needed
func customErrorMessage(e validator.FieldError, FieldName string) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", FieldName)
	default:
		return fmt.Sprintf("Invalid %s", FieldName)
	}
}
func (t *LeadsController) Submit(c echo.Context) error {
	request := models.RequestSubmitLead{}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.GeneralErrorResponse{
			Status:  "error",
			Message: "invalid body request",
		})
	}

	if err := c.Validate(&request); err != nil {
		c.Logger().Error(err.Error())
		ValidationErrors := err.(validator.ValidationErrors)
		for _, ve := range ValidationErrors {
			message := customErrorMessage(ve, ve.Field())
			return c.JSON(http.StatusBadRequest, models.GeneralErrorResponse{
				Status:  "error",
				Message: message,
			})
		}
	}
	submit, err := t.leadService.SubmitLead(request)
	if err != nil {
		errLog := &models.ErrorLogs{
			ErrorMessage: err.Error(),
			Endpoint:     fmt.Sprintf("%s", c.Request().URL),
			StatusCode:   strconv.Itoa(http.StatusBadRequest),
		}
		t.leadRepo.ErrorLogs(errLog)
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.GeneralErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.GeneralSuccessResponse{
		Status: "success",
		ID:     submit.ID,
	})
}

func (t *LeadsController) Leads(c echo.Context) error {
	leads, err := t.leadService.GetLeads()
	if err != nil {
		errLog := &models.ErrorLogs{
			ErrorMessage: err.Error(),
			Endpoint:     fmt.Sprintf("%s", c.Request().URL),
			StatusCode:   strconv.Itoa(http.StatusInternalServerError),
		}
		t.leadRepo.ErrorLogs(errLog)
		return c.JSON(http.StatusInternalServerError, models.GeneralErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, leads)
}

func (t *LeadsController) LeadsByID(c echo.Context) error {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)
	leads, err := t.leadService.GetLeadsByID(idx)
	if err != nil {
		errLog := &models.ErrorLogs{
			ErrorMessage: err.Error(),
			Endpoint:     fmt.Sprintf("%s", c.Request().URL),
			StatusCode:   strconv.Itoa(http.StatusInternalServerError),
		}
		t.leadRepo.ErrorLogs(errLog)
		return c.JSON(http.StatusInternalServerError, models.GeneralErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, leads)
}
