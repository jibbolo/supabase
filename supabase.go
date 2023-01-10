package supabase

import (
	"github.com/labstack/echo/v4"
)

type Supabase struct {
	Auth *Auth
}

func MustNew(projectReference, anonKey string) *Supabase {
	if projectReference == "" || anonKey == "" {
		panic("Missing Supabase projectReference or anonKey")
	}
	return &Supabase{
		NewAuth(projectReference, anonKey),
	}
}

func (s *Supabase) EchoKeyValidation(accessToken string, c echo.Context) (bool, error) {
	user, err := s.Auth.GetUser(accessToken)
	if err != nil {
		return false, err
	}
	c.Set("user", user)
	return true, nil
}
