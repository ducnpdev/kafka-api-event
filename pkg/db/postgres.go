package db

import (
	"encoding/base64"
	"event-tracking/config"
	"event-tracking/pkg/logger"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Username    string `yaml:"username" mapstructure:"username"`
	Password    string `yaml:"password" mapstructure:"password"`
	Database    string `yaml:"database" mapstructure:"database"`
	Host        string `yaml:"host" mapstructure:"host"`
	Port        int    `yaml:"port" mapstructure:"port"`
	Schema      string `yaml:"schema" mapstructure:"schema"`
	MaxIdleConn int    `yaml:"max_idle_conn" mapstructure:"max_idle_conn"`
	MaxOpenConn int    `yaml:"max_open_conn" mapstructure:"max_open_conn"`
}

// create database postgres instance
func InitPostgres(config config.Postgres) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)
	password, _ := base64.StdEncoding.DecodeString(config.Password)
	config.Password = string(password)
	logger.GLogger.Infof("connecting postgres database, user %s, dbname %s, host: %s", config.Username, config.Database, config.Host)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d ", config.Host, config.Username, config.Password, config.Database, config.Port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Default().Println("connect postgres err:", err)
		return db, err
	}
	if config.IsDebug {
		db.Debug()
	}

	dbc, err := db.DB()
	if err == nil {
		if config.MaxIdleConn > 0 {
			dbc.SetMaxIdleConns(config.MaxIdleConn)
		}
		if config.MaxOpenConn > 0 {
			dbc.SetMaxOpenConns(config.MaxOpenConn)
		}
		if config.MaxIdleTime > 0 {
			dbc.SetConnMaxIdleTime(time.Duration(config.MaxIdleTime * int(time.Second)))
		}
		if config.MaxLifeTime > 0 {
			dbc.SetConnMaxLifetime(time.Duration(int(time.Second) * config.MaxLifeTime))
		}
	}

	logger.GLogger.Info("connect postgres successfully")
	return db, err
}
