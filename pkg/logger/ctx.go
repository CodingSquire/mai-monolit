package logger

import (
	"context"
	"github.com/CodingSquire/mai-monolit/pkg/api"

	"github.com/rs/zerolog"


)

func Ctx(ctx context.Context) *api.Logger {
	return zerolog.Ctx(ctx)
}
