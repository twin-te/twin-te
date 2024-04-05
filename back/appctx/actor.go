package appctx

import (
	"context"

	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
)

type actorKey struct{}

func GetActor(ctx context.Context) (authdomain.Actor, bool) {
	actor, ok := ctx.Value(actorKey{}).(authdomain.Actor)
	return actor, ok
}

func SetActor(ctx context.Context, actor authdomain.Actor) context.Context {
	return context.WithValue(ctx, actorKey{}, actor)
}
