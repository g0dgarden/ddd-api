package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"

	"github.com/g0dgarden/ddd-api/utils"
)

// ConfToml　はアプリケーションで使用する設定ファイル
type ConfToml struct {
	SectionDB `toml:"db"`
}

// SectionDB　はDBの設定情報
type SectionDB struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Pass     string `toml:"pass"`
	Database string `toml:"database"`
}

func ConfigureRead(confPath string) (ConfToml, error) {
	var conf ConfToml
	environment := utils.GetEnvironment()
	_, err := toml.DecodeFile(fmt.Sprintf("%s/%s.toml", confPath, environment), &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
