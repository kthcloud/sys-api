package app

import (
	"encoding/json"
	"fmt"
	"sys-api/models"
	"sys-api/pkg/auth"
)

func (context *ClientContext) GetKeycloakToken() (*auth.KeycloakToken, error) {
	tokenRaw, exists := context.GinContext.Get("keycloakToken")
	if !exists {
		return nil, fmt.Errorf("failed to find token in request")
	}

	bytes, err := json.Marshal(tokenRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token in request")
	}

	keycloakToken := auth.KeycloakToken{}
	err = json.Unmarshal(bytes, &keycloakToken)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token in request")
	}

	return &keycloakToken, nil
}

func GetKeycloakConfig() auth.KeycloakConfig {
	var fullCertPath = fmt.Sprintf("realms/%s/protocol/openid-connect/certs", models.Config.Keycloak.Realm)

	return auth.KeycloakConfig{
		Url:           models.Config.Keycloak.URL,
		Realm:         models.Config.Keycloak.Realm,
		FullCertsPath: &fullCertPath,
	}
}
