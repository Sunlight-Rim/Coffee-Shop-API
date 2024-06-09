package model

// easyjson:json
type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    uint64 `json:"phone"`
}

// Signup

// easyjson:json
type DeliverySignupRes struct {
	UserID uint64 `json:"id"`
}

type UsecaseSignupRes struct {
	UserID uint64
}

type StorageCreateRes struct {
	UserID uint64
}

// Signin

type UsecaseSigninRes struct {
	AccessToken  *Token
	RefreshToken *Token
}

type StorageCheckCredentialsRes struct {
	UserID uint64
}

// IsDeleted

type StorageIsDeletedRes struct {
	Deleted bool
}

// Refresh

type UsecaseRefreshRes struct {
	AccessToken  *Token
	RefreshToken *Token
}

// SignoutAll

// easyjson:json
type DeliverySignoutAllRes struct {
	RefreshTokens []string `json:"refresh_tokens"`
}

type UsecaseSignoutAllRes struct {
	RefreshTokens []string
}

// GetMe

// easyjson:json
type DeliveryGetMeRes struct {
	User *User `json:"user"`
}

type UsecaseGetMeRes struct {
	User *User
}

type StorageGetMeRes struct {
	User *User
}

// DeleteMe

// easyjson:json
type DeliveryDeleteMeRes struct {
	User *User `json:"user"`
}

type UsecaseDeleteMeRes struct {
	User *User
}

type StorageDeleteMeRes struct {
	User *User
}
