package rpcs

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/heroiclabs/nakama-common/runtime"
)

type Person struct {
	Name    string   `json:"name" validate:"required"`
	Age     int      `json:"age" validate:"required"`
	Hobbies []string `json:"hobbies" `
	Address Address  `json:"address" validate:"required"`
}

type Address struct {
	HouseNumber int    `json:"houseNumber" validate:"required"`
	Road        string `json:"road" validate:"required"`
}

func PayloadParsing(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	var person Person
	v := validator.New()

	logger.Info("Received payload: %s", payload)

	err := json.Unmarshal([]byte(payload), &person)

	if err != nil {
		logger.Error(err.Error())
		return "", runtime.NewError("Invalid payload shape", INVALID_ARGUMENT)
	}

	err = v.Struct(person)

	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			message := fmt.Sprintf("Payload failed to validate: %+v", validationErrors)
			return "", runtime.NewError(message, INVALID_ARGUMENT)
		} else {
			return "", runtime.NewError("Invalid payload shape", INVALID_ARGUMENT)
		}
	}

	logger.WithField("Person", person).Info("Successfully parsed payload to person struct")

	return "{}", nil
}
