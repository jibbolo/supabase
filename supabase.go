package supabase

import (
	"github.com/labstack/echo/v4"
)

type Supabase struct {
	Anon  *AnonAuth
	Admin *AdminAuth
}

func MustNew(projectReference, anonKey, serviceKey string) *Supabase {
	if projectReference == "" || anonKey == "" {
		panic("Missing Supabase projectReference or anonKey")
	}
	return &Supabase{
		Anon:  &AnonAuth{NewAuth(projectReference, anonKey)},
		Admin: &AdminAuth{NewAuth(projectReference, serviceKey)},
	}
}

func (s *Supabase) EchoKeyValidation(accessToken string, c echo.Context) (bool, error) {
	user, err := s.Anon.GetUser(accessToken)
	if err != nil {
		return false, err
	}
	c.Set("user", user)
	return true, nil
}
