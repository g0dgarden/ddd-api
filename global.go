package api

import (
	"github.com/g0dgarden/ddd-api/conf"
	infra "github.com/g0dgarden/ddd-api/infrastructures"
)

var (
	AppConf conf.ConfToml
	DBConn  infra.Connector
)
