package data

import (
	"context"
	"database/sql"
	"github.com/hayohtee/greenlight/internal/validator"
	"time"
)

// ValidateTokenPlainText check that the plaintext token
// has been provided and is exactly 26 bytes long.
func ValidateTokenPlainText(v *validator.Validator, tokenPlainText string) {
	v.Check(tokenPlainText != "", "token", "must be provided")
	v.Check(len(tokenPlainText) == 26, "token", "must be 26 bytes long")
}

// TokenModel is a struct that wraps a sql.DB connection and methods for
// interacting with tokens table in the database.
type TokenModel struct {
	DB *sql.DB
}

// New creates a new Token struct, inserts the data into the tokens table and return it.
func (m TokenModel) New(userID int64, ttl time.Duration, scope string) (*Token, error) {
	token, err := generateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}
	err = m.Insert(token)
	return token, err
}

// Insert adds the data for a specific token to the tokens table.
func (m TokenModel) Insert(token *Token) error {
	query := `
		INSERT INTO tokens(hash, user_id, expiry, scope)
		VALUES ($1, $2, $3, $4)`

	args := []any{token.Hash, token.UserID, token.Expiry, token.Scope}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, args...)
	return err
}

// DeleteAllForUser deletes all tokens for a specific user and scope.
func (m TokenModel) DeleteAllForUser(scope string, userID int64) error {
	query := `
		DELETE FROM tokens
		WHERE scope = $1 AND user_id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, scope, userID)
	return err
}
