package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type configenv struct {
	PORT        string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
	DB_PORT     string
}

func LoadEnv(env *configenv)(*configenv,error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil,err
	}

	if err := viper.Unmarshal(env); err != nil {
		return nil, err
	}

	return env, nil
}

func Database() (*gorm.DB, error) {
	var env configenv
	ENV,err:=LoadEnv(&env)
	if err!=nil{
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_HOST, ENV.DB_PORT, ENV.DB_DATABASE)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	

	return db, nil
}