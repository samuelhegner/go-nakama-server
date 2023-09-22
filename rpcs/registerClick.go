package rpcs

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/samuelhegner/go-server/helpers"
)

type RegisterClickResponse struct {
	Total int `json:"total"`
}

type ClickRecord struct {
	Count          int `json:"count"`
	ResetTimestamp int `json:"resetTimestamp"`
}

const clickCollection = "ClickCounters"

func RegisterClick(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Info("RegisterClick RPC called")

	userId, err := helpers.GetUserId(ctx, logger)

	if err != nil {
		logger.Error(err.Error())
		return "", runtime.NewError(err.Error(), 13)
	}

	// Read back previous click record if it exists
	record, found, err := helpers.ReadStorageObject(userId, clickCollection, userId, ctx, logger, nk)

	if err != nil {
		logger.WithField("err", err).Error("Storage read error.")
		return "", runtime.NewError(err.Error(), 13)
	}

	var clickRecord ClickRecord

	// If we didn't find the record, we create a new one and write it to storage
	if !found {
		clickRecord = ClickRecord{
			Count:          1,
			ResetTimestamp: 0,
		}

		err := helpers.WriteStorageObject(userId, clickCollection, userId, "*", clickRecord, ctx, logger, nk)

		if err != nil {
			logger.WithField("err", err).Error("Storage write error.")
			return "", runtime.NewError(err.Error(), INTERNAL)
		}

	} else { // if we found the record, we increment the count and write it back
		err := json.Unmarshal([]byte(record.Value), &clickRecord)

		if err != nil {
			logger.WithField("err", err).Error("Storage value unmarshal error.")
			return "", runtime.NewError(err.Error(), INTERNAL)
		}

		clickRecord.Count++

		err = helpers.WriteStorageObject(userId, clickCollection, userId, record.Version, clickRecord, ctx, logger, nk)

		if err != nil {
			logger.WithField("err", err).Error("Storage write error.")
			return "", runtime.NewError(err.Error(), INTERNAL)
		}
	}

	// We return the total click count
	response := &RegisterClickResponse{
		Total: clickRecord.Count,
	}

	resStr, err := helpers.ResponseToJsonString(response, logger)

	return resStr, nil
}
