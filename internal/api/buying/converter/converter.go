package converter

import (
	dtoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/api/buying/dto"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/buying/model"
)

func BuyingDTOToService(dto *dtoModel.Buying) *model.Buying {
	return &model.Buying{
		Name: dto.Name,
	}
}

func BuyingServiceToDTO(model *model.Buying) *dtoModel.Buying {
	return &dtoModel.Buying{
		Name: model.Name,
	}
}