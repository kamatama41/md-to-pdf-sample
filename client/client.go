package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"google.golang.org/api/option"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"

	"golang.org/x/oauth2"
)

var (
	credentialFile = "credentials/credentials.json"
	tokenFile      = "credentials/token.json"
	scopes         = []string{drive.DriveFileScope}
)

func NewService(tok oauth2.TokenSource) (*drive.Service, error) {
	return drive.NewService(context.Background(), option.WithTokenSource(tok))
}

// Request a token from the web, then save the retrieved token and returns it.
func GetTokenFromWeb() (oauth2.TokenSource, error) {
	config, err := getOAuth2Config()
	if err != nil {
		return nil, err
	}

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, err
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, err
	}

	if err := saveToken(tokenFile, tok); err != nil {
		return nil, err
	}

	return config.TokenSource(context.Background(), tok), nil
}

// Retrieves a token from a local file.
func GetTokenFromFile() (oauth2.TokenSource, error) {
	f, err := os.Open(tokenFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)

	config, err := getOAuth2Config()
	if err != nil {
		return nil, err
	}
	return config.TokenSource(context.Background(), tok), err
}

func getOAuth2Config() (*oauth2.Config, error) {
	b, err := ioutil.ReadFile(credentialFile)
	if err != nil {
		return nil, err
	}
	return google.ConfigFromJSON(b, scopes...)
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(token)
}
