package model

// Signup

type SignupRes struct {
	UserID uint64
}

type StorageCreateRes struct {
	UserID uint64
}

// Signin

type SigninRes struct {
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

type RefreshRes struct {
	AccessToken  *Token
	RefreshToken *Token
}

// SignoutAll

type SignoutAllRes struct {
	RefreshTokens []string
}

// GetMe

type GetMeRes struct {
	User *User
}

type StorageGetMeRes struct {
	User *User
}

// DeleteMe

type DeleteMeRes struct {
	User *User
}

type StorageDeleteMeRes struct {
	User *User
}
