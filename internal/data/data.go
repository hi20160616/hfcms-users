package data

import (
	"github.com/hi20160616/hfcms-users/internal/biz"
	"github.com/hi20160616/hfcms-users/internal/data/db/mariadb"
)

var _ biz.UserRepo = new(userRepo)

type Data struct {
	DBClient *mariadb.Client
}
