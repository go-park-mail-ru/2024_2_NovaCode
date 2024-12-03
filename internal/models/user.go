package models

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	UsernameMinLength = 4
	UsernameMaxLength = 16
	PasswordMinLength = 6
	PasswordMaxLength = 32
)

type User struct {
	UserID    uuid.UUID
	Role      string
	Username  string
	Email     string
	Password  string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user.Password = string(hashedPassword)
	return nil
}

func (user *User) ComparePasswords(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (user *User) Validate() error {
	if err := ValidateUsername(user.Username); err != nil {
		return err
	}

	if err := ValidateEmail(user.Email); err != nil {
		return err
	}

	if err := ValidatePassword(user.Password); err != nil {
		return err
	}

	return nil
}

func ValidateUsername(username string) error {
	if len(username) < UsernameMinLength || len(username) > UsernameMaxLength {
		return fmt.Errorf("username length must be at least %d and no more than %d", UsernameMinLength, UsernameMaxLength)
	}

	validChars := regexp.MustCompile(`^[a-zA-Z0-9-_]*$`)
	if !validChars.MatchString(username) {
		return fmt.Errorf("username must contain only latin letters and underscores")

	}

	return nil

}

func ValidateEmail(email string) error {
	emailPattern := regexp.MustCompile(`.+@.+`)
	if !emailPattern.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) < PasswordMinLength || len(password) > PasswordMaxLength {
		return fmt.Errorf("password length must be at least %d and no more than %d", PasswordMinLength, PasswordMaxLength)
	}

	return nil
}
