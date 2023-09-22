package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/samuelhegner/go-server/rpcs"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	initStart := time.Now()

	if err := initializer.RegisterRpc("healthcheck", rpcs.Healthcheck); err != nil {
		return err
	}

	if err := initializer.RegisterRpc("registerClick", rpcs.RegisterClick); err != nil {
		return err
	}

	if err := initializer.RegisterRpc("versionTesting", rpcs.VersionTesting); err != nil {
		return err
	}

	if err := initializer.RegisterRpc("payloadParsing", rpcs.PayloadParsing); err != nil {
		return err
	}

	logger.Info("Module loaded in %dms", time.Since(initStart).Milliseconds())
	return nil
}
