package gozalo

import "fmt"

type DoRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
	AppID        string `json:"app_id"`
	GrantType    string `json:"grant_type"`
}

func (r DoRefreshTokenRequest) ToString() string {
	return fmt.Sprintf("refresh_token=%s&app_id=%s&grant_type=%s", r.RefreshToken, r.AppID, r.GrantType)
}

type DoRefreshTokenResponse struct {
	AccessToken      string `json:"access_token,omitempty"`
	RefreshToken     string `json:"refresh_token,omitempty"`
	ExpiresIn        string `json:"expires_in,omitempty"`
	ErrorName        string `json:"error_name,omitempty"`
	ErrorReason      string `json:"error_reason,omitempty"`
	RefDoc           string `json:"ref_doc,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	Error            int    `json:"error,omitempty"`
}

type GetConversationsResponse struct {
	Data    []MessageDetail `json:"data"`
	Error   int             `json:"error"`
	Message string          `json:"message"`
}

type GetConversationDetailResponse struct {
	Data    []MessageDetail `json:"data"`
	Error   int             `json:"error"`
	Message string          `json:"message"`
}

type SendMessageRequest struct {
	Recipient struct {
		UserId string `json:"user_id"`
	} `json:"recipient"`
	Message struct {
		Attachment struct {
			Payload struct {
				Elements     []AttachmentElement `json:"elements"`
				TemplateType string              `json:"template_type"`
			} `json:"payload"`
			Type string `json:"type"`
		} `json:"attachment"`
		Text string `json:"text"`
	} `json:"message"`
}

type AttachmentElement struct {
	MediaType    string `json:"media_type"`
	AttachmentId string `json:"attachment_id"`
}

type SendMessageResponse struct {
	Data struct {
		MessageID string `json:"message_id"`
		UserID    string `json:"user_id"`
	} `json:"data"`
	Error   int    `json:"error"`
	Message string `json:"message"`
}

type UploadImageResponse struct {
	Data struct {
		AttachmentID string `json:"attachment_id"`
	} `json:"data"`
	Error   int    `json:"error"`
	Message string `json:"message"`
}

type MessageDetail struct {
	Src             int    `json:"src"`
	Time            int64  `json:"time"`
	SentTime        string `json:"sent_time"`
	FromId          string `json:"from_id"`
	FromDisplayName string `json:"from_display_name"`
	FromAvatar      string `json:"from_avatar"`
	ToId            string `json:"to_id"`
	ToDisplayName   string `json:"to_display_name"`
	ToAvatar        string `json:"to_avatar"`
	MessageId       string `json:"message_id"`
	Type            string `json:"type"`
	Url             string `json:"url,omitempty"`
	Message         string `json:"message,omitempty"`
	Thumb           string `json:"thumb,omitempty"`
	Description     string `json:"description,omitempty"`
}

type SendMessageDetailResponse struct {
	Data    MessageDetail `json:"data"`
	Error   int           `json:"error"`
	Message string        `json:"message"`
}
