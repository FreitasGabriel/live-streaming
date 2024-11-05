package service

import (
	"github.com/FreitasGabriel/live-streaming-server/live-streaming-server/internal/model"
	"github.com/FreitasGabriel/live-streaming-server/live-streaming-server/internal/repository"
)

type IKeysService interface {
	AuthStramingKey(name, key string) (*model.Keys, error)
}

type keysService struct {
	repo repository.IKeysRepository
}

func NewKeysService(repo repository.IKeysRepository) IKeysService {
	return &keysService{repo}
}

func (ks *keysService) AuthStramingKey(name, key string) (*model.Keys, error) {

	return ks.repo.FindStreamKey(name, key)
}
