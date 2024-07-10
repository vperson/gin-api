package store

import (
	"database/sql"
	"fmt"
	"gin-api/config"
	log "gin-api/pkg/logger"
	"gin-api/repository/store/tb"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func New(cfg *config.DBConfig) error {
	var err error
	var gd gorm.Dialector
	log.GetSugaredLogger().Infof("store enable: %v", cfg.Enable)
	logStore := log.GetSugaredLogger().With(zap.String("app", "store"))

	if !cfg.Enable {
		return nil
	}

	err = createDatabaseIfNotExist(cfg, logStore)
	if err != nil {
		return err
	}

	gd, err = cfg.CreateGormDialector()
	if err != nil {
		return err
	}

	logging := logger.Default.LogMode(logger.Silent)
	if cfg.Debug {
		logging = logger.Default.LogMode(logger.Info)
	}
	db, err = gorm.Open(gd, &gorm.Config{
		QueryFields: true,
		Logger:      logging,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		return err
	}

	err = autoMigrate()
	if err != nil {
		return err
	}

	err = Common().CreateIfNotExist()
	if err != nil {
		return err
	}
	return nil
}

func autoMigrate() error {
	return db.AutoMigrate(
		&tb.User{},
		&tb.Common{},
	)
}

func createDatabaseIfNotExist(cfg *config.DBConfig, log *zap.SugaredLogger) error {
	if cfg.Tp != config.MySQLDbTp {
		return nil
	}

	// 连接到 MySQL 的默认数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %v", err)
	}
	defer sqlDB.Close()

	// TODO: 数据库不存在直接创建
	_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET %s COLLATE %s;",
		cfg.Database, cfg.Charset, cfg.Collation))
	return err
}
