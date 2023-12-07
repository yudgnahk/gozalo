package gozalo

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	httputils "github.com/yudgnahk/go-common-utils/http"
)

func (c *client) UploadImage(path string) (*UploadImageResponse, error) {
	url := fmt.Sprintf("%s/%s", OpenAPI, UploadImagePath)

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	file, err := os.Open(path)
	defer file.Close()

	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := httputils.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return nil, err
	}

	request.Header.Add("access_token", c.AccessToken)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	var response UploadImageResponse
	err = httputils.Execute(request, &response)

	if response.Error != ErrCodeSuccess {
		return nil, fmt.Errorf("upload image error: %s", response.Message)
	}

	return &response, err
}
