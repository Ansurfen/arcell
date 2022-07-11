package utils

import "github.com/spf13/viper"

func GetConf(confName, confType, dir string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigName(confName)
	conf.SetConfigType(confType)
	conf.AddConfigPath(dir)
	Panic(conf.ReadInConfig())
	return conf
}
