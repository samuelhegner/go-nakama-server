package helpers

import (
	"encoding/json"
	"errors"
	"github.com/heroiclabs/nakama-common/runtime"
)

func ResponseToJsonString(response interface{}, logger runtime.Logger) (string, error) {
	out, err := json.Marshal(response)

	if err != nil {
		logger.Error("Error marshalling response type to JSON: %v", err)
		return "", errors.New("error marshalling response type to JSON")
	}

	return string(out), nil
}
