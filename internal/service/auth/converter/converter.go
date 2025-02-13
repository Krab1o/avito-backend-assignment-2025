package converter

import (
	repoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/user/model"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth/model"
	"golang.org/x/crypto/bcrypt"
)

func CredsServiceToRepo(creds *model.UserCreds) (*repoModel.UserCreds, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(creds.Password), 
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	return &repoModel.UserCreds{
		Username: 		creds.Username,
		PasswordHash: 	string(hashedBytes),
	}, nil
}

func UserServiceToRepo(user *model.User) (*repoModel.User, error) {
	repoCreds, err := CredsServiceToRepo(user.Creds)
	if err != nil {
		return nil, err
	}
	return &repoModel.User{
		Creds: *repoCreds,
		Coins: user.Coins,
	}, nil
}

//TODO: change model to serveModel in all places
// func CredsRepoToService(creds *repoModel.UserCreds) (*model.UserCreds, error) {

// }