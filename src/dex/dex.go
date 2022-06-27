package dex

import "golang.org/x/oauth2"

func getDexConfig() oauth2.Config {
	dex := oauth2.Config{
		ClientID:     "main-application",
		ClientSecret: "rojUkzzGdg6u3RsMzxZykbMcai8jIsen",
		Scopes:       []string{"openid", "email", "groups"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://localhost:5557/auth",
			TokenURL: "http://localhost:5557/token",
		},
	}

	return dex
}
