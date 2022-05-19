package models

import (
	"fmt"
	"log"
	"meta/pkg/setting"
	"meta/pkg/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var Db *gorm.DB

func Setup() {
	var (
		err error
	)

	Db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.MetaDatabaseSetting.User,
		setting.MetaDatabaseSetting.Password,
		setting.MetaDatabaseSetting.Host,
		setting.MetaDatabaseSetting.Name)), &gorm.Config{
		Logger: utils.GormLogger(),
	},
	)

	//user 表使用t591
	Db.Use(dbresolver.Register(dbresolver.Config{
		// `db2` 作为 sources，`db3`、`db4` 作为 replicas
		Sources: []gorm.Dialector{mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.DatabaseSetting.User,
			setting.DatabaseSetting.Password,
			setting.DatabaseSetting.Host,
			setting.DatabaseSetting.Name))},
	}, &User{}))

	if err != nil {
		log.Println(err)
	}

	sqlDB, err := Db.DB()

	if err != nil {
		log.Fatal(2, "can't get DB : %v", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(setting.DatabaseSetting.MaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(setting.DatabaseSetting.MaxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(setting.DatabaseSetting.MaxLifeTime))

	autoErr := Db.AutoMigrate(Question{})
	if autoErr != nil {
		log.Fatalf("AutoMigrate err: %v", autoErr)
	}
}
