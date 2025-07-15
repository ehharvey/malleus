package service

import (
	"github.com/ehharvey/malleus/internal/validation"
)

type ServiceResult[T any] struct {
	Result                  *T
	ModelValidationResult   validation.ModelValidationResult
	DbError                 error
	ServiceValidationResult validation.ServiceValidationResult
}
