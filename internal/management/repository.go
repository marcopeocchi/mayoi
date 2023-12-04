package management

import (
	"os"

	"github.com/marcopeocchi/mayoi/pkg/config"
)

type Repository struct{}

func NewReposityory() *Repository {
	return &Repository{}
}

func (r *Repository) GetDatabaseSize() (int64, error) {
	stat, err := os.Stat(config.Instance().Database)
	if err != nil {
		return 0, err
	}

	return stat.Size(), nil
}
