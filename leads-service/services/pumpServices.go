package services

import (
	"errors"
	"leadsservice/models"
	"leadsservice/repositories"

	"sync"
)

type LeadService interface {
	SubmitLead(request models.RequestSubmitLead) (models.Lead, error)
	GetLeads() ([]models.Lead, error)
	GetLeadsByID(id int) (models.Lead, error)
}

type leadService struct {
	leadRepo repositories.LeadRepository
	mu       sync.Mutex
}

func NewLeadService(leadRepo repositories.LeadRepository) LeadService {
	return &leadService{
		leadRepo: leadRepo,
	}
}

func (s *leadService) SubmitLead(requests models.RequestSubmitLead) (models.Lead, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tx, err := s.leadRepo.BeginTransaction()
	if err != nil {
		return models.Lead{}, err
	}

	emailExist, err := s.leadRepo.GetEmail(requests.Email)
	if err != nil {
		err = s.leadRepo.RollbackTransaction(tx)
		if err != nil {
			return models.Lead{}, err
		}
		return models.Lead{}, err
	}
	if emailExist.Email != "" {
		return models.Lead{}, errors.New("Email already Exist")
	}
	reqData := &models.Lead{
		Name:        requests.Name,
		Email:       requests.Email,
		PhoneNumber: requests.PhoneNumber,
		Source:      requests.Source,
		Message:     requests.Message,
	}

	reqSubmit, err := s.leadRepo.SubmitLead(reqData, tx)
	if err != nil {
		err = s.leadRepo.RollbackTransaction(tx)
		if err != nil {
			return models.Lead{}, err
		}
		return models.Lead{}, err
	}
	err = s.leadRepo.CommitTransaction(tx)
	if err != nil {
		return models.Lead{}, err
	}

	return reqSubmit, nil
}

func (s *leadService) GetLeads() ([]models.Lead, error) {
	leads, err := s.leadRepo.GetLeads()
	if err != nil {
		return nil, err
	}
	return leads, nil
}

func (s *leadService) GetLeadsByID(id int) (models.Lead, error) {
	lead, err := s.leadRepo.GetLeadsByID(id)
	if err != nil {
		return models.Lead{}, err
	}
	return lead, nil
}
