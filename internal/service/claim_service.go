package service

import (
	"claims-backend/internal/model"
	"claims-backend/internal/repository"
	"errors"
	"time"

	"github.com/google/uuid"
)

type ClaimService struct {
	repo repository.ClaimRepository
}

func NewClaimService(repo repository.ClaimRepository) *ClaimService {
	return &ClaimService{
		repo: repo,
	}
}

func (s *ClaimService) CreateClaim(UserID, ClaimType string, amount float64) *model.Claim {
	claim := model.Claim{
		ID:        uuid.New().String(),
		UserID:    UserID,
		Type:      ClaimType,
		Amount:    amount,
		Status:    "SUBMITTED",
		CreatedAt: time.Now(),
	}
	s.repo.Save(claim)
	return &claim
}

func (s *ClaimService) GetClaimByID(id string) (*model.Claim, error) {
	claim, err := s.repo.FindByID(id)
	if err!=nil {
		return nil, err
	}
	return &claim, nil
}
func (s *ClaimService) GetAllClaims() []model.Claim {
	return s.repo.FindAll()
}
func (s *ClaimService) MoveToReview(id string) (*model.Claim, error) {
	claim, err := s.repo.FindByID(id)
	if err!=nil {
		return nil, errors.New("claim not found")
	}
	if claim.Status != model.StatusSubmitted {
		return nil, errors.New("invalid state transition")
	}
	claim.Status = model.StatusUnderReview
	s.repo.Update(claim)
	return &claim, nil
}
func (s *ClaimService) ApproveClaim(id string) (*model.Claim, error) {
	claim, err := s.repo.FindByID(id)
	if err!=nil {
		return nil, errors.New("invalid ID")
	}
	if claim.Status != model.StatusUnderReview {
		return nil, errors.New("only claim under review can be approved")
	}
	claim.Status = model.StatusApproved
	s.repo.Update(claim)
	return &claim, nil
}
func (s *ClaimService) DenyClaim(id string) (*model.Claim, error) {
	claim, err := s.repo.FindByID(id)
	if err!=nil {
		return nil, errors.New("invalid ID")
	}
	if claim.Status != model.StatusUnderReview {
		return nil, errors.New("only claim under review can be denied")
	}
	claim.Status = model.StatusDenied
	s.repo.Update(claim)
	return &claim, nil
}
