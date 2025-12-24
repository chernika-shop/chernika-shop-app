package graph

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// MarshalTime преобразует time.Time в GraphQL scalar
func MarshalTime(t time.Time) graphql.Marshaler {
	return graphql.MarshalTime(t)
}

// UnmarshalTime преобразует GraphQL scalar в time.Time
func UnmarshalTime(v interface{}) (time.Time, error) {
	return graphql.UnmarshalTime(v)
}

// unmarshalInputTime - метод для executionContext (используется generated.go)
func (ec *executionContext) unmarshalInputTime(ctx context.Context, v interface{}) (time.Time, error) {
	return UnmarshalTime(v)
}

// _Time - метод для executionContext (используется generated.go)
func (ec *executionContext) _Time(ctx context.Context, sel ast.SelectionSet, v *time.Time) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return MarshalTime(*v)
}

