package supabase

import (
	"github.com/supabase-community/gotrue-go"
	"github.com/supabase-community/gotrue-go/types"
)

type AuthClient struct {
	gotrue.Client
	ProjectReference,
	ApiKey string
}

func NewAuth(projectReference, apiKey string) *AuthClient {
	return &AuthClient{gotrue.New(
		projectReference,
		apiKey,
	), projectReference, apiKey}
}

type AnonAuth struct {
	client *AuthClient
}

func (a *AnonAuth) GetUser(accessToken string) (User, error) {
	client := a.client.WithToken(accessToken)
	resp, err := client.GetUser()
	if err != nil {
		return User{}, err
	}
	return User{resp.User, accessToken, a}, nil
}

type AdminAuth struct {
	client *AuthClient
}

type MagicLink struct {
	Token    string `json:"token,omitempty"`
	URL      string `json:"url,omitempty"`
	EmailOTP string `json:"email_otp,omitempty"`
}

func (a *AdminAuth) MagicLink(email string) (MagicLink, error) {
	client := a.client.WithToken(a.client.ApiKey)
	resp, err := client.AdminGenerateLink(types.AdminGenerateLinkRequest{
		Type:  types.LinkTypeMagicLink,
		Email: email,
	})
	if err != nil {
		return MagicLink{}, err
	}
	return MagicLink{
		Token:    resp.HashedToken,
		URL:      resp.ActionLink,
		EmailOTP: resp.EmailOTP,
	}, nil
}
