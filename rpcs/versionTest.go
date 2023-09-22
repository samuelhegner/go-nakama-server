package rpcs

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
)

type testStruct struct {
	UpdateCount int `json:"updateCount"`
}

func VersionTesting(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	newDocId := "123456789"
	collection := "script.challenge"
	userId := "00000000-0000-0000-0000-000000000000"
	writeCount := 0
	updateObj := testStruct{
		UpdateCount: 0,
	}

	logger.Info(toJsonString(updateObj))

	var versionId string

	// Write one, version is zeroed
	writeReq := &runtime.StorageWrite{
		Key:        newDocId,
		Collection: collection,
		UserID:     userId,
		Value:      toJsonString(updateObj),
	}

	ack, err := nk.StorageWrite(ctx, []*runtime.StorageWrite{writeReq})

	if err != nil {
		logger.WithField("error", err).Error("Failed to write storage object at write: %d", writeCount)
	}

	versionId = ack[0].Version
	writeCount++
	reportSuccessfulWrite(versionId, writeCount, logger)

	// Write two, version is zeroed
	updateObj.UpdateCount++
	writeReq.Value = toJsonString(updateObj)

	ack, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{writeReq})

	if err != nil {
		logger.WithField("error", err).Error("Failed to write storage object at write: %d", writeCount)
	}

	versionId = ack[0].Version
	writeCount++
	reportSuccessfulWrite(versionId, writeCount, logger)

	// Write three, ack version is provided
	updateObj.UpdateCount++
	writeReq.Value = toJsonString(updateObj)
	writeReq.Version = versionId

	ack, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{writeReq})

	if err != nil {
		logger.WithField("error", err).Error("Failed to write storage object at write: %d", writeCount)
	}

	versionId = ack[0].Version
	writeCount++
	reportSuccessfulWrite(versionId, writeCount, logger)

	// Write four, version is zeroed again
	updateObj.UpdateCount++
	writeReq.Value = toJsonString(updateObj)
	writeReq.Version = ""

	ack, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{writeReq})

	if err != nil {
		logger.WithField("error", err).Error("Failed to write storage object at write: %d", writeCount)
	}

	versionId = ack[0].Version
	writeCount++
	reportSuccessfulWrite(versionId, writeCount, logger)

	return "{}", nil
}

func toJsonString(obj any) string {
	j, _ := json.Marshal(obj)
	return string(j)
}

func reportSuccessfulWrite(version string, writeCount int, logger runtime.Logger) {
	logger.Info("Write %d was successful with version: %s", writeCount, version)
}
