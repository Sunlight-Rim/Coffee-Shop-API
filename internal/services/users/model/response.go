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
type SignupResDelivery struct {
	UserID uint64 `json:"id"`
}

type SignupResUsecase struct {
	UserID uint64
}

type CreateResStorage struct {
	UserID uint64
}

// Signin

type SigninResUsecase struct {
	AccessToken  *Token
	RefreshToken *Token
}

type CheckCredentialsResStorage struct {
	UserID uint64
}

// IsDeleted

type IsDeletedResStorage struct {
	Deleted bool
}

// Refresh

type RefreshResUsecase struct {
	AccessToken  *Token
	RefreshToken *Token
}

// SignoutAll

// easyjson:json
type SignoutAllResDelivery struct {
	RefreshTokens []string `json:"refresh_tokens"`
}

type SignoutAllResUsecase struct {
	RefreshTokens []string
}

// GetMe

// easyjson:json
type GetMeResDelivery struct {
	User *User `json:"user"`
}

type GetMeResUsecase struct {
	User *User
}

type GetMeResStorage struct {
	User *User
}

// DeleteMe

// easyjson:json
type DeleteMeResDelivery struct {
	User *User `json:"user"`
}

type DeleteMeResUsecase struct {
	User *User
}

type DeleteMeResStorage struct {
	User *User
}
