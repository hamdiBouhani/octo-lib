package oauth2

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type Config struct {
	PublicKeyPEM string
	Issuer       string
	Audience     string
}

type Validator struct {
	cfg       Config
	publicKey *rsa.PublicKey
}

func NewValidator(cfg Config) (*Validator, error) {
	block, _ := pem.Decode([]byte(cfg.PublicKeyPEM))
	if block == nil {
		return nil, errors.New("invalid public key PEM")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not RSA public key")
	}

	return &Validator{
		cfg:       cfg,
		publicKey: rsaPub,
	}, nil
}

func (v *Validator) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return v.publicKey, nil
	})
}
