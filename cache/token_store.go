package cache

import (
	"golang.org/x/oauth2"
	"net/http"
)

type TokenStore struct {
	token *oauth2.Token
	state string
}

func (t *TokenStore) StoreToken(token *oauth2.Token, w http.ResponseWriter, req *http.Request) error {
	t.token = token
	return nil
}

func (t *TokenStore) GetToken(req *http.Request) (*oauth2.Token, error) {
	return t.token, nil
}

func (t *TokenStore) StoreState(state string, w http.ResponseWriter, req *http.Request) error {
	t.state = state
	return nil
}

func (t *TokenStore) GetState(*http.Request) (string, error) {
	return t.state, nil
}
