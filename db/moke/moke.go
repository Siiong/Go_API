package moke

import (
	"errors"
	"main/model"

	"main/db"
)

var _ db.StorageSneakers = &Moke{}

type Moke struct {
	listSneakers    map[string]*model.Sneakers `json:"sneakers"`
	listCollections map[string]*model.Collection
}

func New() *db.Storage {
	return &db.Storage{
		Sneakers: &Moke{
			listSneakers:    make(map[string]*model.Sneakers),
			listCollections: make(map[string]*model.Collection),
		},
	}
}

func (m *Moke) GetSneakers(id string) (*model.Sneakers, error) {
	u, ok := m.listSneakers[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return u, nil
}

func (m *Moke) GetAllSneakers() ([]model.Sneakers, error) {
	us := make([]model.Sneakers, len(m.listSneakers))
	var i int
	for k := range m.listSneakers {
		if m.listSneakers[k] != nil {
			us[i] = *m.listSneakers[k]
		}
		i++
	}
	return us, nil
}

func (m *Moke) GetCollection(id string) (*model.Collection, error) {
	u, ok := m.listCollections[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return u, nil
}

func (m *Moke) AddSneakers(s *model.Sneakers) (*model.Sneakers, error) {
	m.listSneakers[s.ID] = s
	return s, nil
}

func (m *Moke) CreateSneakers(s *model.Sneakers) (*model.Sneakers, error) {
	m.listSneakers[s.ID] = s
	return s, nil
}

func (m *Moke) DeleteSneakers(id string) error {
	_, ok := m.listSneakers[id]
	if !ok {
		return errors.New("db user: not found")
	}
	delete(m.listSneakers, id)
	return nil
}
