package gozalo

const (
	OauthURL  = "https://oauth.zaloapp.com/v4/"
	OpenAPI   = "https://openapi.zalo.me/v2.0"
	OpenAPIV3 = "https://openapi.zalo.me/v3.0"

	RefreshTokenPath          = "oa/access_token"
	GetConversationsPath      = "oa/listrecentchat"
	GetConversationDetailPath = "oa/conversation"
	GetUserPath               = "oa/getprofile"

	UploadImagePath = "oa/upload/image"
	SendMessagePath = "oa/message/cs"
)

const (
	ErrCodeSuccess            = 0
	ErrCodeTokenExpired       = -216
	ErrCodeUploadTokenExpired = -155
)
