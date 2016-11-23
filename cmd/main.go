package main

import (
	"flag"

	"github.com/g0dgarden/ddd-api"
	"github.com/g0dgarden/ddd-api/conf"
	infra "github.com/g0dgarden/ddd-api/infrastructures"
	"github.com/g0dgarden/ddd-api/routes"
)

func main() {
	confPath := flag.String("c", "", "configuration file path")
	flag.Parse()

	var err error
	api.AppConf, err = conf.ConfigureRead(*confPath)
	if err != nil {
		panic(err)
	}

	api.DBConn = infra.NewConnection(&api.AppConf.SectionDB)

	router := routers.Init()
	router.Start(":8888")
}
