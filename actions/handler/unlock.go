package handler

import (
	"context"
	"github.com/CTFg/CTFg/actions/gen/action"
	"github.com/CTFg/CTFg/actions/gen/gql"
	"go.uber.org/fx"
)

type unlockDeps struct {
	fx.In

	GQLClient *gql.Client
}

type unlock struct {
	deps unlockDeps
}

func (s *unlock) Unlock(ctx context.Context, args action.UnlockArgs) (response action.UnlockMutationOutput, err error) {

	panic("implement me")
}

func NewUnlock(deps unlockDeps) action.UnlockLogic {
	return &unlock{deps: deps}
}
