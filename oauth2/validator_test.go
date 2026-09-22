package oauth2

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func generateKeyPair(t *testing.T) (*rsa.PrivateKey, string) {
	t.Helper()

	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	pubASN1, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		t.Fatalf("failed to marshal public key: %v", err)
	}

	pemStr := string(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	}))

	return key, pemStr
}

func TestValidator_ValidateToken(t *testing.T) {
	priv, pubPEM := generateKeyPair(t)

	cfg := Config{
		PublicKeyPEM: pubPEM,
	}

	v, err := NewValidator(cfg)
	if err != nil {
		t.Fatalf("validator init failed: %v", err)
	}

	claims := jwt.MapClaims{
		"sub": "user-123",
		"exp": time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signed, err := token.SignedString(priv)
	if err != nil {
		t.Fatalf("failed to sign token: %v", err)
	}

	parsed, err := v.ValidateToken(signed)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !parsed.Valid {
		t.Fatalf("expected token to be valid")
	}
}
