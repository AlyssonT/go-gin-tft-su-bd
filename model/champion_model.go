package model

type ChampionRow struct {
	Id          int
	Name        string
	Tier        int
	Trait       int
	MinToActive int
	IsUnique    bool
}

type Champion struct {
	Name   string `json:"name"`
	Tier   int    `json:"tier"`
	Traits []int  `json:"traits"`
}
