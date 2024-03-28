package service

import (
	"sys-api/dto/body"
	"sys-api/models"
	"sys-api/pkg/repository"
)

func RegisterNode(params *body.HostRegisterParams) error {
	// Validate token
	if params.Token != models.Config.Discovery.Token {
		return BadDiscoveryTokenErr
	}

	if params.DisplayName == "" {
		params.DisplayName = params.Name
	}

	// Register node
	err := repository.NewClient().RegisterHost(models.NewHostByParams(params))
	if err != nil {
		return err
	}

	return nil
}
