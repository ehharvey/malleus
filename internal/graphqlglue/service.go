package graphqlglue

import (
	"github.com/ehharvey/malleus/internal/outcome"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func ServiceResultToGraphql[T any, U any](
	serviceResult outcome.ServiceResult[T],
	mapper func(T) U,
) (U, error) {
	if serviceResult.Succeeded() {
		var result U
		return result, ServiceResultToGraphqlError(serviceResult)
	} else {
		return mapper(serviceResult.Result), nil
	}
}

func ServiceResultToGraphqlError[T any](
	serviceResult outcome.ServiceResult[T],
) error {
	errList := gqlerror.List{}

	return errList
}
