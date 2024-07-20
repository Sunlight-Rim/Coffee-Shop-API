package model

import "time"

type TokensPair struct {
	AccessToken  Token
	RefreshToken Token
}

type Token struct {
	String string
	Exp    time.Time
}
