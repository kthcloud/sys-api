package v2

import (
	"landing-api/pkg/conf"
	"landing-api/pkg/sys"
	"strconv"
)

func GetN(context sys.ClientContext) (int, error) {
	nQuery := context.GinContext.Query("n")
	var n int
	var err error

	if nQuery == "" {
		n = 1
	} else {
		n, _ = strconv.Atoi(nQuery)
		if err != nil {
			n = 1
		}

		if n <= 0 {
			n = 1
		}
	}

	return n, err
}

func IsAdmin(context *sys.ClientContext) bool {
	token, err := context.GetKeycloakToken()
	if err != nil {
		return false
	}

	for _, group := range token.Groups {
		if group == conf.Env.Keycloak.AdminGroup {
			return true
		}
	}

	return false
}
