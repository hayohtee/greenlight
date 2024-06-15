package data

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"
)

// Define constants for the token scope.
const (
	// ScopeActivation represents activation scope
	ScopeActivation = "activation"
	// ScopeAuthentication represents authentication scope.
	ScopeAuthentication = "authentication"
)

// Token is a struct to hold the data for an individual token. This
// includes the plainText and hashed version of the token, associated
// user ID, expiry time and scope.
type Token struct {
	PlainText string    `json:"token"`
	Hash      []byte    `json:"-"`
	UserID    int64     `json:"-"`
	Expiry    time.Time `json:"expiry"`
	Scope     string    `json:"-"`
}

// generateToken returns new token instance containing the given user ID, duration
// the scope.
func generateToken(userId int64, ttl time.Duration, scope string) (*Token, error) {
	token := &Token{
		UserID: userId,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	// Initialize a zero-valued byte slice with a length of 16 bytes.
	randomBytes := make([]byte, 16)

	// Use the Read() function from the crypto/rand package to fill the bytes
	// slice with random bytes from the operating system's CSPRING.
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	// Encode the byte slice to a base-32-encoded string and assign it to the token
	// PlainText field. This will look similar to this:
	//
	// Y3QMGX3PJ3WLRL2YRTQGQ6KRHU
	//
	// The default base-32 strings may be padded at the end with the = character. So
	// we use WithPadding(base32.NoPadding) method to omit them.
	token.PlainText = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	// Generate an SHA-256 hash of the plainText token string.
	// Note that sha256.Sum256() function returns array of length 32, so to make it
	// easier to work with we convert it to a slice using the [:] before storing it.
	hash := sha256.Sum256([]byte(token.PlainText))
	token.Hash = hash[:]

	return token, nil
}
