package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

// ReadStorageObject Returns a storage object, found boolean and error
func ReadStorageObject(key string, collection string, userId string, ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule) (*api.StorageObject, bool, error) {
	objects := []*runtime.StorageRead{{
		Collection: collection,
		Key:        key,
		UserID:     userId,
	},
	}

	records, err := nk.StorageRead(ctx, objects)

	if err != nil {
		logger.WithField("err", err).Error("ReadStorageObject: Storage read error.")
		return nil, false, fmt.Errorf("ReadStorageObject: %w", err)
	}

	if len(records) == 0 {
		return nil, false, nil
	}

	return records[0], true, nil
}

// WriteStorageObject writes the provided object and returns an error if writing fails
func WriteStorageObject(key, collection, userId, version string, value any, ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule) error {
	valueJson, err := json.Marshal(value)

	if err != nil {
		logger.WithField("err", err).Error("Marshal error.")
		return fmt.Errorf("WriteStorageObject: error mashaling value %w", err)
	}

	objectIDs := []*runtime.StorageWrite{{
		Collection: collection,
		Key:        key,
		UserID:     userId,
		Value:      string(valueJson),
	},
	}

	_, writeErr := nk.StorageWrite(ctx, objectIDs)

	if writeErr != nil {
		logger.WithField("err", writeErr).Error("Storage write error.")
		return fmt.Errorf("WriteStorageObject: error writing object %w", err)
	}

	return nil
}
