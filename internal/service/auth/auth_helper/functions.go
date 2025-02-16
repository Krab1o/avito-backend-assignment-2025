package authhelper

import (
	"time"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/shared"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (h *helper) VerifyPassword(hashedPassword string, candidatePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
	return err == nil
}

func (h *helper) GenerateJWT(userID int64, jwtSecret []byte, jwtTimeout int) (string, error) {
    expirationTime := time.Now().Add(time.Duration(jwtTimeout) * time.Minute)
    claims := shared.Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Subject:   "user_auth",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}