package service

import (
	"fmt"

	"github.com/muazwzxv/bookers/m/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

// type DatabaseInterface interface {
// 	Connect() (*gorm.DB, error)
// }

var (
	DB = &Database{}
)

func (orm *Database) Connect() (*gorm.DB, error) {
	config, err := config.DBConfig()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)

	if orm.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return nil, err
	} else {
		orm.DB.Debug().AutoMigrate()
		return orm.DB, nil
	}
}
