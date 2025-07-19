package outcome

type ServiceResult[T any] struct {
	Result                  *T
	Model                   string
	ModelValidationResult   ModelValidationResult
	ServiceValidationResult BusinessValidationResult
	PersistenceResult       DbResult
}

func (sr ServiceResult[T]) Error() string {
	return "TODO"
}
