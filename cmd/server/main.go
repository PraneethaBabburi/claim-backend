package main

import (
	"claims-backend/internal/handler"
	"claims-backend/internal/repository"

	"claims-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	repo := repository.NewClaimRespository()
	claimService := service.NewClaimService(repo)
	claimHandler := handler.NewClaimHandler(claimService)

	r.POST("/claims", claimHandler.CreateClaim)
	r.GET("/claims/:id", claimHandler.GetClaimByID)
	r.GET("/claims", claimHandler.GetAllClaims)

	r.PUT("/claims/:id/review", claimHandler.MoveToReview)
	r.PUT("/claims/:id/approve", claimHandler.ApproveClaim)
	r.PUT("/claims/:id/deny", claimHandler.DenyClaim)

	r.Run(":8080")
}
