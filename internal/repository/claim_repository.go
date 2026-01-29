package repository

import "claims-backend/internal/model"

type ClaimRepository struct {
	data map[string]model.Claim
}

func NewClaimRespository() *ClaimRepository {
	return &ClaimRepository{
		data: make(map[string]model.Claim),
	}
}
func (r *ClaimRepository) Save(claim model.Claim) {
	r.data[claim.ID] = claim
}
func (r *ClaimRepository) FindByID(id string) (model.Claim, bool) {
	claim, exists := r.data[id]
	return claim, exists
}
func (r *ClaimRepository) FindAll() []model.Claim {
	claims := make([]model.Claim, 0, len(r.data))
	for _, claim := range r.data {
		claims = append(claims, claim)
	}
	return claims
}
func (r *ClaimRepository) Update(claim model.Claim) bool {
	if _, exists := r.data[claim.ID]; !exists {
		return false
	}
	r.data[claim.ID] = claim
	return true
}
