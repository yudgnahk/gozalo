package gozalo

import (
	"bytes"
	"fmt"
	"net/http"

	httputils "github.com/yudgnahk/go-common-utils/http"
)

func (c *client) DoRefreshToken() (string, string, error) {
	url := fmt.Sprintf("%s%s", OauthURL, RefreshTokenPath)

	req := &DoRefreshTokenRequest{
		RefreshToken: c.RefreshToken,
		AppID:        c.AppID,
		GrantType:    "refresh_token",
	}

	request, err := httputils.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(req.ToString())))
	if err != nil {
		return "", "", err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Secret_key", c.AppSecret)

	var response DoRefreshTokenResponse
	err = httputils.Execute(request, &response)
	if err != nil {
		return "", "", err
	}

	if response.Error != 0 {
		return "", "", fmt.Errorf("error code: %d, message: %s", response.Error, response.ErrorName)
	}

	return response.AccessToken, response.RefreshToken, nil
}
