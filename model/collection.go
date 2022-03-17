package model

type Collection struct {
	ID   string     `json:"id"`
	Name string     `json:"name"`
	List []Sneakers `json:"list"`
}
