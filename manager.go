package main

import (
	"github.com/heroku/docker-registry-client/registry"
)

type RegistryManager struct {
	hub *registry.Registry
}

func (rm *RegistryManager) Init(url string, username string, password string) (error) {
	rm.hub, err := registry.New(url, username, password)
	return err
}

func (r *RegistryManager) ListRepositories() ([]string, error) {
	return r.hub.Repositories()
}
