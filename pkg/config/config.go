package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const (
	DBName     = "DB_NAME"
	DBUser     = "DB_USER"
	DBPassword = "DB_PASSWORD"
	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
	TimeOut    = "DB_TIMEOUT"
)

func InitConfig() {
	path, _ := os.Getwd()

	path = filepath.Join(path, "..")
	path = filepath.Join(path, "/deploy")

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	//fmt.Println(viper.GetString(DBHost))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Config initialization error: %v", err.Error()))
	}
}
