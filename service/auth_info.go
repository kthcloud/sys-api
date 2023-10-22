package service

import (
	"sys-api/pkg/auth"
	"sys-api/pkg/conf"
)

type AuthInfo struct {
	UserID   string              `json:"userId"`
	JwtToken *auth.KeycloakToken `json:"jwtToken"`
	IsAdmin  bool                `json:"isAdmin"`
}

func CreateAuthInfo(userID string, JwtToken *auth.KeycloakToken, iamGroups []string) *AuthInfo {
	isAdmin := false
	for _, iamGroup := range iamGroups {
		if iamGroup == conf.Env.Keycloak.AdminGroup {
			isAdmin = true
		}
	}

	return &AuthInfo{
		UserID:   userID,
		JwtToken: JwtToken,
		IsAdmin:  isAdmin,
	}
}

func (authInfo *AuthInfo) GetUsername() string {
	return authInfo.JwtToken.PreferredUsername
}

func (authInfo *AuthInfo) GetEmail() string {
	return authInfo.JwtToken.Email
}
