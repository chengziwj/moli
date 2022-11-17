package password

import (
	"errors"
	"fmt"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func Hash(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Verify(hash string, plain string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)); err != nil {
		return false
	}
	return true
}

func CheckPassword(str string) error {
	return Password(str, WithLength(5, 12), WithLower(), WithUpper(), WithNumber(), WithSpecial())
}

type Option func(s string) error

var ErrUpper = errors.New("the password must contain uppercase letters")
var ErrLower = errors.New("the password must contain lowercase letters")
var ErrNumber = errors.New("the password must contain numbers")
var ErrPunctuation = errors.New("the password must contain punctuation")

func WithLength(minLen, maxLen int) Option {
	return func(str string) error {
		if len(str) < minLen || len(str) > maxLen {
			return fmt.Errorf("the password must be %d-%d digits long", minLen, maxLen)
		}
		return nil
	}
}

func WithUpper() Option {
	return func(str string) error {
		for _, s := range str {
			if unicode.IsUpper(s) {
				return nil
			}
		}
		return ErrUpper
	}
}

func WithLower() Option {
	return func(str string) error {
		for _, s := range str {
			if unicode.IsLower(s) {
				return nil
			}
		}
		return ErrLower
	}
}

func WithNumber() Option {
	return func(str string) error {
		for _, s := range str {
			if unicode.IsNumber(s) {
				return nil
			}
		}
		return ErrNumber
	}
}

func WithSpecial() Option {
	return func(str string) error {
		for _, s := range str {
			if unicode.IsSymbol(s) || unicode.IsPunct(s) {
				return nil
			}
		}
		return ErrPunctuation
	}
}

func Password(str string, opts ...Option) error {
	for _, o := range opts {
		if err := o(str); err != nil {
			return err
		}
	}
	return nil
}
