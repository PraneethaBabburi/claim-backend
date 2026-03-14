package repository

import "claims-backend/internal/model"

type ClaimRepository interface {
	Save(claim model.Claim) error
	FindByID(id string) (model.Claim, error)
	FindAll() []model.Claim
	Update(claim model.Claim) bool
}
