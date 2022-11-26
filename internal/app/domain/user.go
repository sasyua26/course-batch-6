package domain

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Signature JWT
var signature = []byte("MyPrivateSignature")

// User struct for db
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	NoHP      string    `json:"no_hp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRegister struct for create new user
type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	NoHP     string `json:"no_hp"`
}

// Function to check input of new user and hash password
func NewUser(email, name, password, hp string) (*User, error) {
	if name == "" {
		return nil, errors.New("Name is required")
	}

	if email == "" {
		return nil, errors.New("Email is required")
	}

	if password == "" {
		return nil, errors.New("Password is required")
	}

	if len(password) < 7 {
		return nil, errors.New("Password length must be 6 characters or more ")
	}

	if hp == "" {
		return nil, errors.New("No. HP is required")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &User{
		Name:     name,
		Email:    email,
		Password: string(hash),
		NoHP:     hp,
	}, nil
}

// Function to generate JWT
func (u User) GenerateJWT() (string, error) {
	claims := jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iss":     "edspert",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(signature)
	return stringToken, err
}

// Function to convert token string to JWT token using signature
func (u User) DecryptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}
		return signature, nil
	})

	data := make(map[string]interface{})
	if err != nil {
		return data, err
	}

	if !parsedToken.Valid {
		return data, errors.New("Invalid token")
	}

	return parsedToken.Claims.(jwt.MapClaims), nil
}
