package model

type Trait struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	MinToActivate int    `json:"minToActivate"`
}
