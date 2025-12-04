package utility

import (
	"os"
	"strings"
)

func GetSecrets(key string, defaultValue string) string {
	fileKey := key + "_FILE"

	if secretFile := os.Getenv(fileKey); secretFile != "" {
		content, err := os.ReadFile(secretFile)

		if err == nil{
			return strings.TrimSpace(string(content))
		}
	}

	return defaultValue
}