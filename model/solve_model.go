package model

type DataToSolve struct {
	Champions map[int]Champion
	Traits    map[int]int
}

type Solution struct {
	Champions  []Champion  `json:"champions"`
	Traits     map[int]int `json:"traits"`
	Evaluation int         `json:"evaluation"`
}
