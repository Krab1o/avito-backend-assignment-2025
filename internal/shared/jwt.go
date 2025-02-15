package shared

import "github.com/golang-jwt/jwt/v5"

const UserIDJsonName = "user_id"

type Claims struct {
    UserID int64 `json:"user_id"`
    jwt.RegisteredClaims
}
