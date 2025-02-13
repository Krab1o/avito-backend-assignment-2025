package converter

import (
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/auth/dto"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth/model"
)

func CredsDTOToService(creds *dto.UserCreds) *model.UserCreds {
	return &model.UserCreds{
		Username: creds.Username,
		Password: creds.Password,
	}
}

func CredsServiceToDTO(creds *model.UserCreds) *dto.UserCreds {
	return &dto.UserCreds{
		Username: creds.Username,
		Password: creds.Password,
	}
}