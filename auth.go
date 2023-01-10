package supabase

import (
	"github.com/supabase-community/gotrue-go"
)

type Auth struct {
	gotrue.Client
	ProjectReference,
	AnonKey string
}

func NewAuth(projectReference, anonKey string) *Auth {
	return &Auth{gotrue.New(
		projectReference,
		anonKey,
	), projectReference, anonKey}
}

func (a *Auth) GetUser(accessToken string) (User, error) {
	client := a.WithToken(accessToken)
	resp, err := client.GetUser()
	if err != nil {
		return User{}, err
	}
	return User{resp.User, accessToken, a}, nil
}
