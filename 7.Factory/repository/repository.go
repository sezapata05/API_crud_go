package repository

import (
	"factory/configuration"
	"factory/repository/mysql"
	"factory/repository/sqlserver"
	"fmt"
)

type Repository interface {
	Find(id int) string
	Save(data string) error
}

func New(config *configuration.Configuration) (Repository, error) {
	var repo Repository
	var err error

	switch config.Engine {
	case "mysql":
		repo = mysql.New()
	case "sqlserver":
		repo = sqlserver.New()
	default:
		fmt.Errorf("invalid engine %s", config.Engine)

	}

	return repo, err
}
