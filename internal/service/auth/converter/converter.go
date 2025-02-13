package converter

import (
	repoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/user/model"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth/model"
)

const baseID = 0

//TODO: add hashing to passwords
func CredsServiceToRepo(creds *model.UserCreds) *repoModel.UserCreds {
	return &repoModel.UserCreds{
		Username: 		creds.Username,
		PasswordHash: 	creds.Password,
	}
}