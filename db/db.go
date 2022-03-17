package db

import "main/model"

type Storage struct {
	Sneakers StorageSneakers
}

type StorageSneakers interface {
	GetAllSneakers() ([]model.Sneakers, error)
	GetSneakers(id string) (*model.Sneakers, error)
	/* GetAllCollection() ([]model.Collection, error)
	GetCollection(id string) (*model.Collection, error) */
	CreateSneakers(s *model.Sneakers) (*model.Sneakers, error)
	/*CreateCollection() (*model.Collection, error)
	AddSneakers(id_collection string, id_sneakers string) error */
	DeleteSneakers(id_sneakers string) error
}
