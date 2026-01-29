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
	r.Run(":8080")
}
