package models

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              uuid.UUID `json:"id"`
	CreatedAt       time.Time `json:"_"`
	UpdatedAt       time.Time `json:"_"`
	Email           string    `json:"email"`
	PasswordHash    string    `json:"-"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:password_confirm`
}

func (u *User) Register(conn *pgx.Conn) error {

	u.Email = strings.ToLower(u.Email)
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("erro creating account")
	}

	u.PasswordHash = string(pwdHash)
	now := time.Now()

	_, err = conn.Exec(context.Background(), "INSERT INTO user_account(created_at, updated_at,email, password_hash) VALUES($1, $2, $3, $4)", now, now, u.Email, u.PasswordHash)

	return err

}

func (u *User) GetAuthToken() (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	authToken, err := token.SignedString("DGFSGDFG$%#%WERFDVFGSGDFDS")
	fmt.Print("Got token" + authToken)
	return authToken, err

}
