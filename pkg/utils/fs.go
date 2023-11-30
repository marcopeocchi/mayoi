package utils

import (
	"os"

	"github.com/marcopeocchi/mayoi/pkg/config"
)

func GetDatabaseSize() (int64, error) {
	stat, err := os.Stat(config.Instance().Database)
	if err != nil {
		return 0, err
	}

	return stat.Size(), nil
}
