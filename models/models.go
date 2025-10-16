package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type RegisterRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
}

type RelationShipRequest struct {
	Inviter string `json:"inviter"`
	Invitee string `json:"invitee"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshToken struct {
}

type Tokens struct {
	Auth_token    string
	Refresh_token string
}

type FavorRequest struct {
	Creator_id   string `json:"creator"`
	Recipient_id string `json:"recipient"`
	Favor_text   string `json:"favortext"`
}

type AuthClaims struct {
	UserName string `json:"userName"`
	UserId   string `json:"userId"`
	jwt.RegisteredClaims
}
