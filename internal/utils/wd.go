package openuem_utils

import (
	"os"
	"path/filepath"
)

func GetWd() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}
