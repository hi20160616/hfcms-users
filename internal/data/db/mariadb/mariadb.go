package mariadb

import (
	"database/sql"

	"github.com/hi20160616/hfcms-users/configs"
)

type Client struct {
	db             *sql.DB
	DatabaseClient *DatabaseClient
}

type DatabaseClient struct {
	db *sql.DB
}

func open(cfg *configs.Config) (*sql.DB, error) {
	return sql.Open(cfg.Database.Driver, cfg.Database.Source)
}

func NewClient(projectName configs.ProjectName) (*Client, error) {
	cfg := configs.NewConfig(projectName)
	if cfg.Err != nil {
		return &Client{nil, nil}, cfg.Err
	}
	db, err := open(cfg)
	return &Client{db, &DatabaseClient{db}}, err
}
