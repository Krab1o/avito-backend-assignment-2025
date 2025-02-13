package converter

import (
	repoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/inventory/model"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/buying/model"
)

func BuyingServiceToRepo(buying *model.Buying) *repoModel.Buying {
	return &repoModel.Buying{
		Name: buying.Name,
	}
}

func BuyingRepoToService(buying *repoModel.Buying) *model.Buying {
	return &model.Buying{
		Name: buying.Name,
	}
}