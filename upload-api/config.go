package upload_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func getStoresConfig() ([]string, error) {

	storesConfigString := os.Getenv("SYMBOL_STORE_LOCAL_STORES")
	if storesConfigString == "" {
		return nil, errors.New("No local stores configured")
	}

	var storesConfig []string
	if err := json.Unmarshal([]byte(storesConfigString), &storesConfig); err != nil {
		return nil, errors.New(fmt.Sprintf("Error while decoding local stores configuration: %w", err))
	}

	return storesConfig, nil
}
