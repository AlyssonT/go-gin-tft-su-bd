package usecase

import (
	"tft_su_bd_backend/model"
	"tft_su_bd_backend/repository"
)

type TraitUsecase struct {
	repository repository.TraitRepository
}

func NewTraitUsecase(repo repository.TraitRepository) TraitUsecase {
	return TraitUsecase{
		repository: repo,
	}
}

func (pu *TraitUsecase) GetTraits() ([]model.Trait, error) {
	return pu.repository.GetTraits()
}
