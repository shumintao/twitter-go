package internal

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:",omitempty"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (u *User) prepareForInsert() error {
	password := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	refreshToken := uuid.New().String()
	u.Password = string(hashedPassword)
	u.RefreshToken = refreshToken
	return nil
}

func (u *User) sanitize() {
	u.Password = ""
}

func (u *User) compareHashAndPassword(password string) error {
	p := []byte(password)
	hp := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(hp, p)
}

type AuthorizeDto struct {
	Username string `json:"username"`
	Password string `json:",omitempty"`
}

type AuthorizeResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}