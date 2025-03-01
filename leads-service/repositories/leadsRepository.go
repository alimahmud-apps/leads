package repositories

import (
	"database/sql"
	"errors"
	"leadsservice/config"
	"leadsservice/models"
	"log"
)

type LeadRepository interface {
	SubmitLead(request *models.Lead, tx *sql.Tx) (models.Lead, error)
	GetEmail(email string) (models.Lead, error)
	GetLeads() ([]models.Lead, error)
	ErrorLogs(data *models.ErrorLogs) (models.ErrorLogs, error)
	GetLeadsByID(id int) (models.Lead, error)

	BeginTransaction() (*sql.Tx, error)
	CommitTransaction(*sql.Tx) error
	RollbackTransaction(*sql.Tx) error
}

type leadRepo struct {
}

func (r *leadRepo) BeginTransaction() (*sql.Tx, error) {
	return config.DB.Begin()
}
func (r *leadRepo) CommitTransaction(tx *sql.Tx) error {
	return tx.Commit()
}
func (r *leadRepo) RollbackTransaction(tx *sql.Tx) error {
	return tx.Rollback()
}
func NewLeadsRepository() LeadRepository {
	return &leadRepo{}
}

func (r *leadRepo) SubmitLead(request *models.Lead, tx *sql.Tx) (models.Lead, error) {
	var insertedID int
	query := `INSERT INTO leads (name,email,phone_number,source,message) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	err := tx.QueryRow(query, request.Name, request.Email, request.PhoneNumber, request.Source, request.Message).Scan(&insertedID)
	if err != nil {
		return models.Lead{}, err
	}
	request.ID = insertedID
	return *request, nil
}

func (r *leadRepo) GetEmail(email string) (models.Lead, error) {
	query := `SELECT id,name,email,phone_number,source,message FROM leads where email=$1`
	var lead models.Lead
	err := config.DB.Get(&lead, query, email)
	if err != nil {
		if err == sql.ErrNoRows { // Check if no rows are returned
			return models.Lead{}, nil
		}
		log.Println("Failed to execute query:", err)
		return models.Lead{}, err
	}

	return lead, nil
}

func (r *leadRepo) GetLeads() ([]models.Lead, error) {
	query := `SELECT id,name,email,phone_number,source,message,created_at FROM leads ORDER BY ID DESC`

	var leads []models.Lead
	err := config.DB.Select(&leads, query)
	if err != nil {
		log.Println("Failed to execute query:", err)
		return nil, err
	}
	// Handle no results
	if len(leads) == 0 {
		return []models.Lead{}, nil
	}
	return leads, nil
}

func (r *leadRepo) ErrorLogs(data *models.ErrorLogs) (models.ErrorLogs, error) {
	var insertedID int
	query := `INSERT INTO error_logs (error_message, endpoint, status_code) VALUES ($1,$2,$3) RETURNING id`
	err := config.DB.QueryRow(query, data.ErrorMessage, data.Endpoint, data.StatusCode).Scan(&insertedID)
	if err != nil {
		return models.ErrorLogs{}, err
	}
	data.ID = insertedID
	return *data, nil
}

func (r *leadRepo) GetLeadsByID(id int) (models.Lead, error) {
	query := `SELECT id,name,email,phone_number,source,message,created_at FROM leads where id=$1`

	var lead models.Lead
	err := config.DB.Get(&lead, query, id)
	if err != nil {
		if err == sql.ErrNoRows { // Check if no rows are returned
			return models.Lead{}, errors.New("no data")
		}
		log.Println("Failed to execute query:", err)
		return models.Lead{}, err
	}
	return lead, nil
}
