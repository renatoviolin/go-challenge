package test_utils

import (
	"api-unico/infra/logger"

	"github.com/subosito/gotenv"
)

func LoadVars() {
	paths := []string{"./.env", "../.env", "../../.env", "../../../.env", "../../../../.env", "../../../../../.env"}
	for _, p := range paths {
		err := gotenv.Load(p)
		if err == nil {
			return
		}
	}
	logger.LogFatal("load-vars", "unable to find .env file")
}
