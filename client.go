package gozalo

// Client is a client for the Zalo OA API.
type client struct {
	AppID        string
	AppSecret    string
	AccessToken  string
	RefreshToken string
}

type Client interface {
	DoRefreshToken() (string, string, error)
	GetConversations(count, offset int) (GetConversationsResponse, error)
	GetConversationDetail(userID string, count, offset int) (GetConversationDetailResponse, error)
	SendMessage(userID string, message string, attachmentID *string) (SendMessageDetailResponse, error)
	UploadImage(path string) (*UploadImageResponse, error)
}

// NewClient returns a new Zalo OA API client.
func NewClient(appID, appSecret, accessToken, refreshToken string) Client {
	return &client{
		AppID:        appID,
		AppSecret:    appSecret,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
