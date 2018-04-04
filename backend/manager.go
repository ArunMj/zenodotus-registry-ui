package main

import (
	"github.com/rusenask/docker-registry-client/registry"
)

type RegistryManager struct {
	hub *registry.Registry
}

func CreateManager(url string, username string, password string) (*RegistryManager, error) {
	hub, err := registry.New(url, username, password)
	if err != nil {
		return nil, err
	}
	return &RegistryManager{hub: hub}, nil
}

func (r *RegistryManager) ListRepositories() ([]string, error) {
	return r.hub.Repositories()
}
