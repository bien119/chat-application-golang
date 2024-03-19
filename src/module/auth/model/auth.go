package model

import "time"

type Auth struct {
	Id            string     `json:"id" bson:"_id,omitempty"`
	Username      string     `json:"username" bson:"username"`
	Email         string     `json:"email" bson:"email"`
	Role          string     `json:"role" bson:"role"`
	EmailVerified bool       `json:"email_verified" bson:"emailVerified"`
	CreatedAt     *time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt     *time.Time `json:"updated_at" bson:"updatedAt"`
}

type AuthEmailPassword struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Token struct {
	Token string `json:"token"`
	// ExpiredIn in seconds
	ExpiredIn int64 `json:"expire_in"`
}

type TokenResponse struct {
	AccessToken Token `json:"access_token"`
	// RefreshToken will be used when access token expired
	// to issue new pair access token and refresh token.
	RefreshToken *Token `json:"refresh_token,omitempty"`
}

func (Auth) TableName() string { return "users" }