package supabase

import (
	"github.com/supabase-community/gotrue-go"
	"github.com/supabase-community/gotrue-go/types"
)

type Auth struct {
	gotrue.Client
	ProjectReference,
	ApiKey string
}

func NewAuth(projectReference, apiKey string) *Auth {
	return &Auth{gotrue.New(
		projectReference,
		apiKey,
	), projectReference, apiKey}
}

type AnonAuth struct {
	*Auth
}

func (a *AnonAuth) GetUser(accessToken string) (User, error) {
	client := a.WithToken(accessToken)
	resp, err := client.GetUser()
	if err != nil {
		return User{}, err
	}
	return User{resp.User, accessToken, a}, nil
}

type AdminAuth struct {
	*Auth
}

type MagicLink struct {
	Token string
	URL   string
}

func (a *AdminAuth) MagicLink(email string) (MagicLink, error) {
	client := a.WithToken(a.ApiKey)
	resp, err := client.AdminGenerateLink(types.AdminGenerateLinkRequest{
		Type:  types.LinkTypeMagicLink,
		Email: email,
	})
	if err != nil {
		return MagicLink{}, err
	}
	return MagicLink{resp.HashedToken, resp.ActionLink}, nil
}
