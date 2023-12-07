package gozalo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	httputils "github.com/yudgnahk/go-common-utils/http"
)

func (c *client) GetConversations(count, offset int) (GetConversationsResponse, error) {
	var conversations GetConversationsResponse

	data := url.QueryEscape(fmt.Sprintf("{\"offset\":%d,\"count\":%d}", offset, count))
	reqUrl := fmt.Sprintf("%s/%s?data=%v", OpenAPI, GetConversationsPath, data)

	request, err := httputils.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return conversations, err
	}

	request.Header.Add("access_token", c.AccessToken)

	err = httputils.Execute(request, &conversations)
	if err != nil {
		return conversations, err
	}

	return conversations, err
}

func (c *client) GetConversationDetail(userID string, count, offset int) (GetConversationDetailResponse, error) {
	var conversationDetail GetConversationDetailResponse

	data := url.QueryEscape(fmt.Sprintf("{\"offset\":%d,\"user_id\":%s,\"count\":%d}", offset, userID, count))
	reqUrl := fmt.Sprintf("%s/%s?data=%v", OpenAPI, GetConversationDetailPath, data)

	request, err := httputils.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return conversationDetail, err
	}

	request.Header.Add("access_token", c.AccessToken)

	err = httputils.Execute(request, &conversationDetail)
	if err != nil {
		return conversationDetail, err
	}

	return conversationDetail, err
}

func (c *client) SendMessage(userID string, message string, attachmentID *string) (SendMessageDetailResponse, error) {
	reqUrl := fmt.Sprintf("%s/%s", OpenAPIV3, SendMessagePath)

	var req SendMessageRequest
	req.Recipient.UserId = userID

	if attachmentID != nil {
		req.Message.Attachment.Payload.Elements = []AttachmentElement{
			{
				AttachmentId: *attachmentID,
				MediaType:    "image",
			},
		}

		req.Message.Attachment.Payload.TemplateType = "media"
		req.Message.Attachment.Type = "template"
	} else {
		req.Message.Text = message
	}

	// marshal request data
	data, err := json.Marshal(req)
	if err != nil {
		return SendMessageDetailResponse{}, fmt.Errorf("error marshall request: %w", err)
	}

	request, err := httputils.NewRequest(http.MethodPost, reqUrl, bytes.NewBuffer(data))
	if err != nil {
		return SendMessageDetailResponse{}, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("access_token", c.AccessToken)

	var response SendMessageResponse
	err = httputils.Execute(request, &response)
	if err != nil {
		return SendMessageDetailResponse{}, err
	}

	if response.Error != ErrCodeSuccess {
		return SendMessageDetailResponse{
			Error:   response.Error,
			Message: response.Message,
		}, nil
	}

	// get latest message of user_id
	conversationDetail, err := c.GetConversationDetail(userID, 1, 0)
	if err != nil {
		return SendMessageDetailResponse{}, err
	}

	if conversationDetail.Error != ErrCodeSuccess {
		return SendMessageDetailResponse{
			Error:   conversationDetail.Error,
			Message: conversationDetail.Message,
		}, nil
	}

	return SendMessageDetailResponse{
		Data: conversationDetail.Data[0],
	}, nil
}
