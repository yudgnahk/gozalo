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
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
}

type GetConversationsResponse struct {
	Data []struct {
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
		Thumb           string `json:"thumb,omitempty"`
		Url             string `json:"url,omitempty"`
		Description     string `json:"description,omitempty"`
		Message         string `json:"message,omitempty"`
	} `json:"data"`
	Error   int    `json:"error"`
	Message string `json:"message"`
}

type GetConversationDetailResponse struct {
	Data []struct {
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
	} `json:"data"`
	Error   int    `json:"error"`
	Message string `json:"message"`
}

type SendMessageRequest struct {
	Recipient struct {
		UserId string `json:"user_id"`
	} `json:"recipient"`
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
}

type SendMessageResponse struct {
	Data struct {
		MessageID string `json:"message_id"`
		UserID    string `json:"user_id"`
	} `json:"data"`
	Error   int    `json:"error"`
	Message string `json:"message"`
}
