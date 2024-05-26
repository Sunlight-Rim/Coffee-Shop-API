package rest

// easyjson:json
type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    uint64 `json:"phone"`
}

// easyjson:json
type SignupRes struct {
	UserID uint64 `json:"id"`
}

// easyjson:json
type SignoutAllRes struct {
	RefreshTokens []string `json:"refresh_tokens"`
}

// easyjson:json
type GetMeRes struct {
	User User `json:"user"`
}

// easyjson:json
type DeleteMeRes struct {
	User User `json:"user"`
}
