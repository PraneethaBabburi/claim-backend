package main

import (
	"claims-backend/internal/handler"
	"claims-backend/internal/repository"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"claims-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := sql.Open(
		"mysql",
		"root:Praneetha@327@tcp(127.0.0.1:3306)/claims_db?parseTime=true",
	)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	repo := repository.NewMySQLClaimRepository(db)
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
