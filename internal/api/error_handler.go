package api

import (
	"errors"
	"log"
	"net/http"

	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

func HandleServiceError(c *gin.Context, err error) {
	var appErr *errs.AppError
	check := errors.As(err, &appErr)
	switch appErr.ErrType {
		//send response
	case errs.ServiceError:
		c.JSON(http.StatusInternalServerError, gin.H{FieldError: appErr.Message})
	case errs.NotFound:
		c.JSON(http.StatusNotFound, gin.H{FieldError: appErr.Message})
	case errs.SemanticError:
		c.JSON(http.StatusBadRequest, gin.H{FieldError: appErr.Message})
	case errs.Unauthorized:
		c.JSON(http.StatusUnauthorized, gin.H{FieldError: appErr.Message})
	}
	for check {
		log.Printf("[%s]: %s", appErr.ErrType, appErr.Message)
		check = errors.As(appErr.Err, &appErr)
	}
}