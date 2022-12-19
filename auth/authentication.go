package auth

import (
	"TradingBot/cache"
	"context"
	"fmt"
	"github.com/zricethezav/go-tdameritrade"
	"golang.org/x/oauth2"
	"net/http"
	"os"
)

type Authentication struct {
	authenticator *tdameritrade.Authenticator
}

func NewAuthentication() *Authentication {
	return &Authentication{
		authenticator: tdameritrade.NewAuthenticator(&cache.TokenStore{},
			oauth2.Config{
				ClientID: os.Getenv("TDA_CONSUMER_KEY"),
				Endpoint: oauth2.Endpoint{
					TokenURL: "https://api.tdameritrade.com/v1/oauth2/token",
					AuthURL:  os.Getenv("TDA_AUTH_URL"),
				},
				RedirectURL: os.Getenv("TDA_REDIRECT_URL"),
			}),
	}
}

func (a *Authentication) Callback(w http.ResponseWriter, req *http.Request) error {
	fmt.Println("Processing TDA Callback...")

	_, err := a.authenticator.FinishOAuth2Flow(context.Background(), w, req)

	if err != nil {
		fmt.Println("Error processing TDA callback!", err)
		return err
	}

	return nil
}

func (a *Authentication) Authenticate(w http.ResponseWriter, req *http.Request) (string, error) {

	fmt.Println("Authenticating with TDA...")

	redirectURl, err := a.authenticator.StartOAuth2Flow(w, req)

	if err != nil {
		fmt.Println("Error occurred authenticating", err)
		return "", err
	}

	fmt.Println("Redirect URL: ", redirectURl)

	return redirectURl, nil
}

func (a *Authentication) GetAuthenticatedClient() (*tdameritrade.Client, error) {
	client, err := a.authenticator.AuthenticatedClient(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (a *Authentication) IsLoggedIn(r *http.Request) (bool, error) {

	token, err := a.authenticator.Store.GetToken(r)

	if err != nil {
		fmt.Println("Error retrieving token", err)
		return false, err
	}

	if token != nil {
		return true, nil
	}

	return false, nil
}
