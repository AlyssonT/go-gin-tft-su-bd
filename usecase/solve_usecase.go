package usecase

import (
	"fmt"
	"tft_su_bd_backend/model"
	"tft_su_bd_backend/repository"
)

type SolveUseCase struct {
	repository repository.SolveRepository
}

func NewSolveUseCase(repo repository.SolveRepository) SolveUseCase {
	return SolveUseCase{
		repository: repo,
	}
}

func (pu *SolveUseCase) GetDataToSolve(augment string) (model.DataToSolve, error) {
	championRows, err := pu.repository.GetChampionsWithTraits()

	if err != nil {
		fmt.Println(err)
		return model.DataToSolve{}, err
	}

	var dataToSolve model.DataToSolve
	dataToSolve.Champions = make(map[int]model.Champion)
	dataToSolve.Traits = make(map[int]int)
	for _, championRow := range championRows {
		champion, exist := dataToSolve.Champions[championRow.Id]
		if !championRow.IsUnique || augment == "builtDifferent" {
			if exist {
				champion.Traits = append(champion.Traits, championRow.Trait)
				dataToSolve.Champions[championRow.Id] = champion

			} else {
				dataToSolve.Champions[championRow.Id] = model.Champion{
					Name:   championRow.Name,
					Tier:   championRow.Tier,
					Traits: []int{championRow.Trait},
				}
			}
		}

		_, exist = dataToSolve.Traits[championRow.Trait]
		if !exist {
			dataToSolve.Traits[championRow.Trait] = championRow.MinToActive
		}
	}

	return dataToSolve, nil
}
