package gozalo

import (
	"fmt"
	"net/url"

	httputils "github.com/yudgnahk/go-common-utils/http"
)

func (c *client) GetUser(id string) (GetUserResponse, error) {
	var user GetUserResponse

	data := url.QueryEscape(fmt.Sprintf("{\"user_id\":%s}", id))

	reqUrl := fmt.Sprintf("%s/%s?data=%v", OpenAPI, GetUserPath, data)

	request, err := httputils.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return user, err
	}

	request.Header.Add("access_token", c.AccessToken)

	err = httputils.Execute(request, &user)
	if err != nil {
		return user, err
	}

	return user, err
}
