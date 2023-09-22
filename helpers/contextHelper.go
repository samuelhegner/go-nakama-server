package helpers

import (
	"context"
	"errors"
	"github.com/heroiclabs/nakama-common/runtime"
)

func GetUserId(ctx context.Context, logger runtime.Logger) (string, error) {

	userId := ctx.Value(runtime.RUNTIME_CTX_USER_ID)

	if userId == nil {
		logger.Error("unable to retrieve UserId from provided context")
		return "", errors.New("unable to retrieve UserId from provided context")
	}

	return userId.(string), nil
}
