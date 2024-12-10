package repository

import (
	"database/sql"
	"fmt"
	"tft_su_bd_backend/model"
)

type TraitRepository struct {
	connection *sql.DB
}

func NewTraitRepository(connection *sql.DB) TraitRepository {
	return TraitRepository{
		connection,
	}
}

func (pr *TraitRepository) GetTraits() ([]model.Trait, error) {
	query := "SELECT id, name, minToActivate FROM trait"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Trait{}, err
	}

	defer rows.Close()

	var traitList []model.Trait
	var traitObj model.Trait

	for rows.Next() {
		err = rows.Scan(&traitObj.Id, &traitObj.Name, &traitObj.MinToActivate)

		if err != nil {
			fmt.Println(err)
			return []model.Trait{}, err
		}

		traitList = append(traitList, traitObj)
	}

	return traitList, nil
}
