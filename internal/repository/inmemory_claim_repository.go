package repository

import "claims-backend/internal/model"

type InMemoryClaimRepository struct {
	data map[string]model.Claim
}

func NewInMemoryClaimRepository() *InMemoryClaimRepository {
	return &InMemoryClaimRepository{
		data: make(map[string]model.Claim),
	}
}
func (r *InMemoryClaimRepository) Save(claim model.Claim) {
	r.data[claim.ID] = claim
}
func (r *InMemoryClaimRepository) FindByID(id string) (model.Claim, bool) {
	claim, exists := r.data[id]
	return claim, exists
}
func (r *InMemoryClaimRepository) FindAll() []model.Claim {
	claims := make([]model.Claim, 0, len(r.data))
	for _, claim := range r.data {
		claims = append(claims, claim)
	}
	return claims
}
func (r *InMemoryClaimRepository) Update(claim model.Claim) bool {
	if _, exists := r.data[claim.ID]; !exists {
		return false
	}
	r.data[claim.ID] = claim
	return true
}
