package logger

import (
	"context"
	"gitlab.com/CodingSquire/mai_monolit/pkg/api"

	"github.com/rs/zerolog"


)

func Ctx(ctx context.Context) *api.Logger {
	return zerolog.Ctx(ctx)
}
