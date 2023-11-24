package config

import (
	"github.com/spf13/viper"
	"log"
)

var cfgReader *configReader

type (
	Configuration struct {
		DatabaseSettings
		JwtSettings
	}
	// DatabaseSettings 数据库配置
	DatabaseSettings struct {
		DatabaseURL  string
		DatabaseName string
		Username     string
		Password     string
	}
	// JwtSettings JWT配置
	JwtSettings struct {
		SecretKey string
	}
	configReader struct {
		configFile string
		v          *viper.Viper
	}
)

// 实例化reader
func newConfigReader(configFile string) {
	v := viper.GetViper()
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)
	cfgReader = &configReader{
		configFile: configFile,
		v:          v,
	}
}

// GetAllConfigValues 获取配置
func GetAllConfigValues(configFile string) (configuration *Configuration, err error) {
	// 实例化reader
	newConfigReader(configFile)
	if err = cfgReader.v.ReadInConfig(); err != nil {
		log.Printf("配置文件读取失败：%v", err)
		return nil, err
	}
	err = cfgReader.v.Unmarshal(&configuration)
	if err != nil {
		log.Printf("配置文件解析失败：%v", err)
		return nil, err
	}
	return configuration, err
}
