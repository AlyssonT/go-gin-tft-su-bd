package repository

import (
	"database/sql"
	"fmt"
	"tft_su_bd_backend/model"
)

type SolveRepository struct {
	connection *sql.DB
}

func NewSolveRepository(connection *sql.DB) SolveRepository {
	return SolveRepository{
		connection,
	}
}

func (pr *SolveRepository) GetChampionsWithTraits() ([]model.ChampionRow, error) {
	query := "SELECT c.id, c.name, c.tier, t.id, t.mintoactivate, t.isunique FROM champion c INNER JOIN championtraits ct ON c.id = ct.id_champion INNER JOIN trait t ON t.id = ct.id_trait"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.ChampionRow{}, err
	}

	defer rows.Close()

	var championRowList []model.ChampionRow
	var championRowObj model.ChampionRow

	for rows.Next() {
		err = rows.Scan(&championRowObj.Id, &championRowObj.Name, &championRowObj.Tier, &championRowObj.Trait, &championRowObj.MinToActive, &championRowObj.IsUnique)

		if err != nil {
			fmt.Println(err)
			return []model.ChampionRow{}, err
		}

		championRowList = append(championRowList, championRowObj)
	}

	return championRowList, nil
}
