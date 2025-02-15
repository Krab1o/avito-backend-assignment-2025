package auth

import (
	"context"
	"log"
	"time"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth/converter"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	shared "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/jwt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const startCoins = 1000

func generateJWT(userID int64, jwtSecret []byte, jwtTimeout int) (string, error) {
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

func verifyPassword(hashedPassword string, candidatePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
	return err == nil
}

// Concerns separation violation for the sake of efficiency
func (s *serv) Auth(ctx context.Context, userCreds *model.UserCreds) (string, error) {
	var jwtUserID int64
	repoData, err := s.userRepository.GetUserByUsername(ctx, nil, userCreds.Username)
	if err != nil {
		return "", errs.NewServiceError(service.MessageInternalError, err)
	}
	if repoData == nil {
		newUser := &model.User{
			Creds: userCreds,
			Coins: startCoins,
		}
		newUserRepo, err := converter.UserServiceToRepo(newUser)
		if err != nil {
			return "", errs.NewServiceError(service.MessageInternalError, err)
		}
		jwtUserID, err = s.userRepository.CreateUser(ctx, nil, newUserRepo)
		if err != nil {
			return "", errs.NewServiceError(service.MessageInternalError, err)
		}
	} else {
		ok := verifyPassword(
			repoData.Creds.PasswordHash,
			userCreds.Password,
		)
		if !ok {
			return "", errs.NewUnauthorizedError(service.MessageWrongPassword, err)
		}
		jwtUserID = repoData.ID
	}
	log.Println(s.jwtConfig.Timeout())
	token, err := generateJWT(jwtUserID, s.jwtConfig.Secret(), s.jwtConfig.Timeout())
	if err != nil {
		return "", errs.NewServiceError(service.MessageInternalError, err)
	}
	return token, nil
}
