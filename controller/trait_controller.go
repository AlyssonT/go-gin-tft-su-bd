package controller

import (
	"net/http"
	"tft_su_bd_backend/model"
	"tft_su_bd_backend/usecase"

	"github.com/gin-gonic/gin"
)

type TraitController struct {
	usecase usecase.TraitUsecase
}

func NewTraitController(usecase usecase.TraitUsecase) TraitController {
	return TraitController{
		usecase,
	}
}

func (pc *TraitController) GetTraits(ctx *gin.Context) {
	traits, err := pc.usecase.GetTraits()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response[*struct{}]{
			Status: http.StatusInternalServerError,
			Data:   nil,
		})
	}

	ctx.JSON(http.StatusOK, model.Response[[]model.Trait]{
		Status: http.StatusOK,
		Data:   traits,
	})
}
