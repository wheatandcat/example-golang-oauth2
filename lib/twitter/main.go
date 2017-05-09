package twitter

import (
	"encoding/json"

	"errors"

	"github.com/astaxie/beego"
	"github.com/garyburd/go-oauth/oauth"
)

var (
	tempCredKey  string
	tokenCredKey string
)

// Account アカウント
type Account struct {
	ID         string `json:"id_str"`
	ScreenName string `json:"screen_name"`
}

func init() {
	tempCredKey = beego.AppConfig.String("twitterConsumerKey")
	tokenCredKey = beego.AppConfig.String("twitterConsumerSecret")
}

// GetConnect 接続を取得する
func GetConnect() *oauth.Client {
	return &oauth.Client{
		TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
		TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
		Credentials: oauth.Credentials{
			Token:  tempCredKey,
			Secret: tokenCredKey,
		},
	}
}

// GetAccessToken アクセストークンを取得する
func GetAccessToken(rt *oauth.Credentials, oauthVerifier string) (*oauth.Credentials, error) {
	oc := GetConnect()
	at, _, err := oc.RequestToken(nil, rt, oauthVerifier)

	return at, err
}

// GetMe 自身を取得する
func GetMe(at *oauth.Credentials, user *Account) error {
	oc := GetConnect()
	resp, err := oc.Get(nil, at, "https://api.twitter.com/1.1/account/verify_credentials.json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return errors.New("Twitter is unavailable")
	}

	if resp.StatusCode >= 400 {
		return errors.New("Twitter request is invalid")
	}

	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return err
	}

	return nil

}
