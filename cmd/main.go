package main

import (
	"tft_su_bd_backend/controller"
	"tft_su_bd_backend/db"
	"tft_su_bd_backend/repository"
	"tft_su_bd_backend/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada repository
	TraitRepository := repository.NewTraitRepository(dbConnection)
	SolveRepository := repository.NewSolveRepository(dbConnection)

	// Camada usecase
	TraitUseCase := usecase.NewTraitUsecase(TraitRepository)
	SolveUseCase := usecase.NewSolveUseCase(SolveRepository)

	// Camada de controllers
	TraitController := controller.NewTraitController(TraitUseCase)
	SolveController := controller.NewSolveController(SolveUseCase)

	server.GET("/traits", TraitController.GetTraits)
	server.GET("/solve/:numChamps", SolveController.GetSolution)
	server.Run(":8000")
}
