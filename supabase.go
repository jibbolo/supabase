package supabase

import (
	"github.com/labstack/echo/v4"
)

type Supabase struct {
	Anon  *AnonAuth
	Admin *AdminAuth
}

func MustNew(projectReference, anonKey, serviceKey string) *Supabase {
	if projectReference == "" || anonKey == "" || serviceKey == "" {
		panic("Missing Supabase projectReference, anonKey or serviceKey")
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
