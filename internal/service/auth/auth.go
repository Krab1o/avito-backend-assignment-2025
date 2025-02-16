package auth

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth/converter"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
)

const StartCoins = 1000

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
			Coins: StartCoins,
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
		ok := s.authHelper.VerifyPassword(
			repoData.Creds.PasswordHash,
			userCreds.Password,
		)
		if !ok {
			return "", errs.NewUnauthorizedError(service.MessageWrongPassword, err)
		}
		jwtUserID = repoData.ID
	}
	token, err := s.authHelper.GenerateJWT(jwtUserID, s.jwtConfig.Secret(), s.jwtConfig.Timeout())
	if err != nil {
		return "", errs.NewServiceError(service.MessageInternalError, err)
	}
	return token, nil
}
