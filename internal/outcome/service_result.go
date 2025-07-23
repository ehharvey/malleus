package outcome

import "errors"

type ServiceResult[T any] struct {
	Result                  T
	Model                   string
	ModelValidationResult   ModelValidationResult
	ServiceValidationResult BusinessValidationResult
	PersistenceResult       DbResult
}

func (sr ServiceResult[T]) CombineErrors() error {
	return errors.Join(
		sr.ModelValidationResult.CombineErrors(),
		sr.ServiceValidationResult.CombineErrors(),
		sr.PersistenceResult,
	)
}

func (sr ServiceResult[T]) Succeeded() bool {
	return sr.ModelValidationResult.Succeeded() &&
		sr.ServiceValidationResult.Succeded() &&
		sr.PersistenceResult.Succeeded()
}
