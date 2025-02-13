package auth

import (
	"context"
	"errors"
	"time"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth/converter"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth/model"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const startCoins = 1000

// Concerns separation violation for the sake of efficiency
func (s *serv) Auth(ctx context.Context, userCreds *model.UserCreds) (string, error) {
	user := &model.User{
		Creds: userCreds,
	}
	userRepo, err := converter.UserServiceToRepo(user)
	if err != nil {
		return "", err
	}
	repoCredsData, err := s.userRepository.FindUserCreds(ctx, &userRepo.Creds)
	if err != nil {
		return "", err
	}
	if repoCredsData == nil {
		user.ID, err = s.userRepository.CreateUser(ctx, userRepo)
	} else {
		err = bcrypt.CompareHashAndPassword(
			[]byte(userRepo.Creds.PasswordHash),
			[]byte(user.Creds.Password),
		)
		if err != nil {
			return "", errors.New("Wrong password")
		}
	}
	return GenerateJWT(user.ID, s.jwtConfig.Secret())
}

type Claims struct {
    UserID int64 `json:"user_id"`
    jwt.RegisteredClaims
}

func GenerateJWT(userID int64, jwtSecret string) (string, error) {
    expirationTime := time.Now().Add(5 * time.Minute) // Срок жизни Access-токена
    
    claims := Claims{
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