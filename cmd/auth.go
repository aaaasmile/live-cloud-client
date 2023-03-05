package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"live-cloud-client/conf"
	"net/http"
)

type Token struct {
	AccessToken   string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
	TokenType     string `json:"token_type"`
	Expire        string `json:"expiry"`
	RefreshExpire string `json:"refresh_expiry"`
}

type CredentialResponse struct {
	Info       string
	ResultCode int
	Username   string
	Token      Token
}

type CredentialReq struct {
	Username string
	Password string
	UserHash string
}

func getReqToken() (string, error) {
	url := conf.Current.ServiceURL + "Token"
	cr := CredentialReq{
		Username: conf.Current.Username,
		UserHash: conf.Current.UserHash,
	}
	b, err := json.Marshal(cr)
	if err != nil {
		return "", err
	}
	reader := bytes.NewReader(b)
	resp, err := http.Post(url, "application/json", reader)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	rawbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error on token request: %s", string(rawbody))
	}
	creResp := CredentialResponse{}
	if err := json.Unmarshal(rawbody, &creResp); err != nil {
		return "", err
	}

	tk := creResp.Token.AccessToken

	return tk, nil
}
