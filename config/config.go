package config

import (
	"fmt"
	"sync"
)

// AppConfig Application configuration
type AppConfig struct {
	Port int
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

// GetConfig Initiatilize config in singleton way
func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8080

	fmt.Println(defaultConfig)

	return &defaultConfig
}
