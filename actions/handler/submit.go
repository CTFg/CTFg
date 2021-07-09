package handler

import (
	"context"
	"github.com/CTFg/CTFg/actions/gen/action"
	"github.com/CTFg/CTFg/actions/gen/gql"
	"go.uber.org/fx"
)

type submitDeps struct {
	fx.In

	GQLClient *gql.Client
}

type submit struct {
	deps submitDeps
}

func (s *submit) Submit(ctx context.Context, args action.SubmitArgs) (response action.SubmitAttemptMutationOutput, err error) {

	panic("implement me")
}

func NewSubmit(deps submitDeps) action.SubmitLogic {
	return &submit{deps: deps}
}
