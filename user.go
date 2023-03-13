package supabase

import (
	"fmt"

	"github.com/jibbolo/postgrest-go"
	"github.com/supabase-community/gotrue-go/types"
	storage_go "github.com/supabase-community/storage-go"
)

type User struct {
	types.User
	AccessToken string
	auth        *AnonAuth
}

func (u *User) NewAuthenticatedRestClient() *postgrest.Client {
	url := fmt.Sprintf("https://%s.supabase.co/rest/v1/", u.auth.ProjectReference)
	client := postgrest.NewClient(url, "", nil)
	client.ApiKey(u.auth.ApiKey)
	client.TokenAuth(u.AccessToken)
	return client
}

func (u *User) NewAuthenticatedStorageClient() *storage_go.Client {
	url := fmt.Sprintf("https://%s.supabase.co/storage/v1", u.auth.ProjectReference)
	return storage_go.NewClient(url, u.AccessToken, nil)
}
