package handler

import (
	"claims-backend/internal/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClaimHandler struct {
	claimService *service.ClaimService
}

func NewClaimHandler(cs *service.ClaimService) *ClaimHandler {
	return &ClaimHandler{claimService: cs}
}

type CreateClaimRequest struct {
	UserID    string  `json:"user_id"`
	ClaimType string  `json:"type"`
	Amount    float64 `json:"amount"`
}

func (h *ClaimHandler) CreateClaim(c *gin.Context) {
	var req CreateClaimRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	claim := h.claimService.CreateClaim(req.UserID, req.ClaimType, req.Amount)
	c.JSON(http.StatusCreated, claim)
}

func (h *ClaimHandler) GetClaimByID(c *gin.Context) {
	id := c.Param("id")
	claim, exists := h.claimService.GetClaimByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.New("Invalid ID")})
		return
	}
	c.JSON(http.StatusOK, claim)
}

func (h *ClaimHandler) GetAllClaims(c *gin.Context) {
	claim := h.claimService.GetAllClaims()
	c.JSON(http.StatusOK, claim)
}
func (h *ClaimHandler) MoveToReview(c *gin.Context) {
	id := c.Param("id")
	claim, err := h.claimService.MoveToReview(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, claim)

}
func (h *ClaimHandler) ApproveClaim(c *gin.Context) {
	id := c.Param("id")
	claim, err := h.claimService.ApproveClaim(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, claim)
}
func (h *ClaimHandler) DenyClaim(c *gin.Context) {
	id := c.Param("id")
	claim, err := h.claimService.DenyClaim(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, claim)
}
